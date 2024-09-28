package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/hankeyyh/chat-box-svr/conf"
	"github.com/hankeyyh/chat-box-svr/dao"
	"github.com/hankeyyh/chat-box-svr/model"
	"github.com/sashabaranov/go-openai"
)

func AppPublicList(req *http.Request) (interface{}, *zerror) {
	apps, err := dao.App.AllPublic()
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return apps, nil
}

func AppPrivateList(req *http.Request) (interface{}, *zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user_id"), 10, 64)
	if userId == 0 || err != nil{
		return nil, NewZError(-1, "user_id is required", err)
	}
	// TODO validate string format
	apps, err := dao.App.AllPrivateByAuthor(userId)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return apps, nil
}

func AppDetail(req *http.Request) (interface{}, *zerror) {
	appIDStr := req.Form.Get("app_id")
	appId, err := strconv.ParseUint(appIDStr, 10, 64)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	app, err := dao.App.GetByID(appId)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}

	return app, nil
}

func AppUpsert(req *http.Request) (interface{}, *zerror) {
	request := AppUpsertRequest{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	app := model.App{
		Id:              request.AppId,
		ModelId:         request.ModelId,
		Name:            request.Name,
		Temperature:     request.Temperature,
		TopP:            request.TopP,
		MaxOutputTokens: request.MaxOutputTokens,
		Context:         request.Context,
		CreatedBy:       request.CreatedBy,
		Introduction:    request.Introduction,
		Prologue:        request.Prologue,
		Prompt:          request.Prompt,
		IsPublic:        request.IsPublic,
	}
	err = dao.App.Save(&app)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	data := AppUpsertResponseData{
		AppId: app.Id,
	}
	return data, nil
}

func AppRelease(req *http.Request) (interface{}, *zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user_id"), 10, 64)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	request := AppReleaseRequest{}
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	// check if app is created by user
	app, err := dao.App.GetByAuthorAndId(userId, request.AppId)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	app.IsPublic = 1
	if err = dao.App.Save(&app); err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return nil, nil
}

func AppUnrelease(req *http.Request) (interface{}, *zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user_id"), 10, 64)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	request := AppUnReleaseRequest{}
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	// check if app is created by user
	app, err := dao.App.GetByAuthorAndId(userId, request.AppId)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	app.IsPublic = 0
	if err = dao.App.Save(&app); err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	
	return nil, nil
}

func AppChatList(req *http.Request) (interface{}, *zerror) {
	appId, err := strconv.ParseUint(req.Form.Get("app_id"), 10, 64)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	var chatId uint64
	if req.Form.Get("chat_id") != "" {
		var err error
		chatId, err = strconv.ParseUint(req.Form.Get("chat_id"), 10, 64)
		if err != nil {
			return nil, NewZError(-1, err.Error(), err)
		}
	}
	page, err := strconv.Atoi(req.Form.Get("page"))
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	pageSize, err := strconv.Atoi(req.Form.Get("page_size"))
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	userId, err := strconv.ParseUint(req.Header.Get("user_id"), 10, 64)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	offset := (page - 1) * pageSize
	historyList, err := dao.ChatHistory.BatchGetRecentByUserID(appId, userId, chatId, offset, pageSize)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return historyList, nil
}

// TODO openaikey should send from client
func AppChat(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Content-Type must be application/json"))
		return
	}
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Method must be POST"))
		return
	}

	chatReq := AppChatRequest{}
	err := json.NewDecoder(req.Body).Decode(&chatReq)
	if err != nil {
		returnError(w, err)
		return
	}
	appId := chatReq.AppId
	content := chatReq.Content

	userId, err := strconv.ParseUint(req.Header.Get("user_id"), 10, 64)
	if err != nil {
		returnError(w, err)
		return
	}

	app, err := dao.App.GetByID(appId)
	if err != nil {
		returnError(w, err)
		return
	}
	aimodel, err := dao.AiModel.GetByID(app.ModelId)
	if err != nil {
		returnError(w, err)
		return
	}

	// save to chat list
	userChat := model.ChatHistory{
		AppId:    app.Id,
		ParentId: nil,
		UserId:   userId,
		Sender: "user",
		Content: content,
	}
	if err = dao.ChatHistory.Save(&userChat); err != nil {
		returnError(w, err)
		return
	}

	serverConf := conf.DefaultConf.ServerConf
	openaiConf := openai.DefaultConfig(serverConf.Key)
	openaiConf.BaseURL = serverConf.BaseUrl
	client := openai.NewClientWithConfig(openaiConf)
	openaiReq := openai.ChatCompletionRequest{
		Model:       aimodel.Name,
		Temperature: app.Temperature,
		TopP:        app.TopP,
		MaxTokens:   int(app.MaxOutputTokens),
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: app.Prompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(context.Background(), openaiReq)
	if err != nil {
		returnError(w, err)
		return
	}
	defer stream.Close()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.(http.Flusher).Flush()

	// write stream to buffer
	var assistantBuf bytes.Buffer
	for {
		openaiRsp, err := stream.Recv()
		if err != nil && !errors.Is(err, io.EOF) {
			returnError(w, err)
			w.(http.Flusher).Flush()
			return
		}
		end := errors.Is(err, io.EOF)
		content := ""
		if !end {
			content = openaiRsp.Choices[0].Delta.Content
		}
		assistantBuf.WriteString(content)
		rsp := Response{
			Code:    0,
			Message: "",
			Data: ChatSpan{
				Content: content,
				End:     end,
			},
		}
		json.NewEncoder(w).Encode(rsp)
		w.(http.Flusher).Flush()
		if end {
			break
		}
	}
	// save assistant chat
	assistantChat := model.ChatHistory{
		AppId:    app.Id,
		ParentId: &userChat.Id,
		UserId:   userId,
		Sender: "assistant",
		Content: assistantBuf.String(),
	}
	if err = dao.ChatHistory.Save(&assistantChat); err != nil {
		returnError(w, err)
		return
	}
}

func returnError(w http.ResponseWriter, err error) {
	rsp := Response{
		Code:    -1,
		Message: err.Error(),
		Data:    nil,
	}
	json.NewEncoder(w).Encode(rsp)
}

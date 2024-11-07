package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/sashabaranov/go-openai"

	"github.com/hankeyyh/chat-box-svr/conf"
	"github.com/hankeyyh/chat-box-svr/constant"
	"github.com/hankeyyh/chat-box-svr/dao"
	"github.com/hankeyyh/chat-box-svr/model"
	"github.com/hankeyyh/chat-box-svr/util/log"
	"github.com/hankeyyh/chat-box-svr/zerror"
)

func SessionList(req *http.Request) (interface{}, zerror.Zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user-id"), 10, 64)
	if userId == 0 || err != nil {
		return nil, zerror.NewZError(-1, "user-id is required", err)
	}
	sessions, err := dao.Session.GetByUserID(userId)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	return sessions, nil
}

func SessionCreate(req *http.Request) (interface{}, zerror.Zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user-id"), 10, 64)
	if userId == 0 || err != nil {
		return nil, zerror.NewZError(-1, "user-id is required", err)
	}
	var sessionCreateReq SessionCreateRequest
	if err = json.NewDecoder(req.Body).Decode(&sessionCreateReq); err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	// 创建session
	session := model.Session{
		UserId: userId,
		Name:   sessionCreateReq.Name,
	}
	if err = dao.Session.Create(&session); err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	return session, nil
}

func SessionUpdate(req *http.Request) (interface{}, zerror.Zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user-id"), 10, 64)
	if userId == 0 || err != nil {
		return nil, zerror.NewZError(-1, "user-id is required", err)
	}
	var sessionUpdateReq SessionUpdateRequest
	if err = json.NewDecoder(req.Body).Decode(&sessionUpdateReq); err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	if sessionUpdateReq.Id == 0 {
		return nil, zerror.NewZError(-1, "session_id is required", nil)
	}
	// 检查session是否属于用户
	session, err := dao.Session.GetByID(sessionUpdateReq.Id)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	if session.UserId != userId {
		return nil, zerror.NewZError(-1, "session not belongs to user", nil)
	}
	// 更新
	session.Name = sessionUpdateReq.Name
	if err = dao.Session.Save(&session); err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	return session, nil
}

func SessionChatList(req *http.Request) (interface{}, zerror.Zerror) {
	var chatId uint64
	if req.Form.Get("chat_id") != "" {
		var err error
		chatId, err = strconv.ParseUint(req.Form.Get("chat_id"), 10, 64)
		if err != nil {
			return nil, zerror.NewZError(-1, err.Error(), err)
		}
	}
	page, err := strconv.Atoi(req.Form.Get("page"))
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	pageSize, err := strconv.Atoi(req.Form.Get("page_size"))
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	sessionId, err := strconv.ParseUint(req.Form.Get("session_id"), 10, 64)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	offset := (page - 1) * pageSize
	historyList, err := dao.ChatHistory.BatchGetRecentBySessionID(sessionId, chatId, offset, pageSize)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	return historyList, nil
}

// assistant stream message
func SessionStream(w http.ResponseWriter, req *http.Request) {
	sseBuf := &SSEBuffer{}
	req.ParseForm()
	userId, err := strconv.ParseUint(req.Header.Get("user-id"), 10, 64)
	if err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
	aiKey := req.Header.Get("ai-key")
	if aiKey == "" {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: "ai-key is required", Data: nil}))
		return
	}
	// parse request
	chatId, err := strconv.ParseUint(req.Form.Get("chat_id"), 10, 64)
	if err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
	// load chat
	chat, err := dao.ChatHistory.GetByID(chatId)
	if err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
	if chat.UserId != userId {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: "chat not belongs to user", Data: nil}))
		return
	}
	// load session
	session, err := dao.Session.GetByID(chat.SessionId)
	if err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
	// load app
	app, err := dao.App.GetByID(chat.AppId)
	if err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
	// load model
	aimodel, err := dao.AiModel.GetByID(app.ModelId)
	if err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
	// build openai request
	serverConf := conf.DefaultConf.ServerConf
	openaiConf := openai.DefaultConfig(aiKey)
	openaiConf.BaseURL = serverConf.BaseUrl
	client := openai.NewClientWithConfig(openaiConf)
	openaiReq, err := buildChatCompletionRequest(aimodel, app, session)
	if err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
	stream, err := client.CreateChatCompletionStream(context.Background(), *openaiReq)
	if err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
	defer stream.Close()

	// write stream header
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.(http.Flusher).Flush()

	// write stream rsp to buffer
	var assistantBuf bytes.Buffer
	for {
		openaiRsp, err := stream.Recv()
		if err != nil && !errors.Is(err, io.EOF) {
			w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
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
		log.Infof("response: %+v", rsp)
		w.Write(sseBuf.EncodeJson(rsp))
		if end {
			break
		}
	}
	// save assistant chat
	assistantChat := model.ChatHistory{
		AppId:     app.Id,
		SessionId: session.Id,
		ParentId:  &chat.Id,
		UserId:    userId,
		Sender:    "assistant",
		Content:   assistantBuf.String(),
	}
	if err = dao.ChatHistory.Save(&assistantChat); err != nil {
		w.Write(sseBuf.EncodeJson(Response{Code: -1, Message: err.Error(), Data: nil}))
		return
	}
}

func SessionChat(req *http.Request) (interface{}, zerror.Zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user-id"), 10, 64)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}

	chatReq := SessionChatRequest{}
	if err = json.NewDecoder(req.Body).Decode(&chatReq); err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	sessionId := chatReq.SessionId
	appId := chatReq.AppId
	content := chatReq.Content
	if appId == 0 {
		appId = conf.DefaultConf.ChatConf.DefaultAppId
	}
	// if no session, create one
	var session model.Session
	if sessionId == 0 {
		session = model.Session{
			UserId: userId,
		}
		if err = dao.Session.Create(&session); err != nil {
			return nil, zerror.NewZError(-1, err.Error(), err)
		}
	} else {
		// check if session belongs to user
		session, err = dao.Session.GetByID(sessionId)
		if err != nil {
			return nil, zerror.NewZError(-1, err.Error(), err)
		}
		if session.UserId != userId {
			return nil, zerror.NewZError(-1, "session not belongs to user", nil)
		}
	}

	app, err := dao.App.GetByID(appId)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	// save to chat list
	userChat := model.ChatHistory{
		AppId:     app.Id,
		SessionId: session.Id,
		ParentId:  nil,
		UserId:    userId,
		Sender:    "user",
		Content:   content,
	}
	if err = dao.ChatHistory.Save(&userChat); err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}

	return SessionChatResponseData{ChatId: userChat.Id}, nil
}

type SSEBuffer struct {
	bytes.Buffer
}

func (sb *SSEBuffer) EncodeJson(v interface{}) []byte {
	sb.WriteString("data: ")
	json.NewEncoder(sb).Encode(v)
	// json.Encode 会在末尾添加\n，这里再添加一个\n，以符合SSE消息格式
	sb.WriteByte('\n')
	result := sb.Bytes()
	sb.Reset()
	return result
}

func buildChatCompletionRequest(aiModel model.AiModel, app model.App, session model.Session) (*openai.ChatCompletionRequest, error) {
	openaiReq := openai.ChatCompletionRequest{
		Model:       aiModel.Name,
		Temperature: app.Temperature,
		TopP:        app.TopP,
		MaxTokens:   int(app.MaxOutputTokens),
		Stream:      true,
	}
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: app.Prompt,
		},
	}
	// load chat history
	chatHistoryList, err := dao.ChatHistory.GetAllBySessionID(session.Id)
	if err != nil {
		return nil, err
	}

	for _, chatHistory := range chatHistoryList {
		role := openai.ChatMessageRoleUser
		if chatHistory.Sender == constant.RoleAssistant {
			role = openai.ChatMessageRoleAssistant
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    role,
			Content: chatHistory.Content,
		})
	}

	openaiReq.Messages = messages

	return &openaiReq, nil
}

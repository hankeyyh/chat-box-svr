package logic

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hankeyyh/chat-box-svr/dao"
	"github.com/hankeyyh/chat-box-svr/model"
	"github.com/hankeyyh/chat-box-svr/zerror"
)

func AppPublicList(req *http.Request) (interface{}, zerror.Zerror) {
	apps, err := dao.App.AllPublic()
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	return apps, nil
}

func AppPrivateList(req *http.Request) (interface{}, zerror.Zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user-id"), 10, 64)
	if userId == 0 || err != nil {
		return nil, zerror.NewZError(-1, "user-id is required", err)
	}
	// TODO validate string format
	apps, err := dao.App.AllPrivateByAuthor(userId)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	return apps, nil
}

func AppDetail(req *http.Request) (interface{}, zerror.Zerror) {
	appIDStr := req.Form.Get("app_id")
	appId, err := strconv.ParseUint(appIDStr, 10, 64)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	app, err := dao.App.GetByID(appId)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}

	return app, nil
}

func AppUpsert(req *http.Request) (interface{}, zerror.Zerror) {
	request := AppUpsertRequest{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	app := model.App{
		Id:              request.AppId,
		ModelId:         request.ModelId,
		Name:            request.Name,
		Temperature:     request.Temperature,
		TopP:            request.TopP,
		MaxOutputTokens: request.MaxOutputTokens,
		CreatedBy:       request.CreatedBy,
		Introduction:    request.Introduction,
		Prologue:        request.Prologue,
		Prompt:          request.Prompt,
		IsPublic:        request.IsPublic,
		ShowPrompt: 	request.ShowPrompt,
		Icon:            request.Icon,
	}
	err = dao.App.Save(&app)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	data := AppUpsertResponseData{
		AppId: app.Id,
	}
	return data, nil
}

func AppRelease(req *http.Request) (interface{}, zerror.Zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user-id"), 10, 64)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	request := AppReleaseRequest{}
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	// check if app is created by user
	app, err := dao.App.GetByAuthorAndId(userId, request.AppId)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	app.IsPublic = 1
	if err = dao.App.Save(&app); err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	return nil, nil
}

func AppUnrelease(req *http.Request) (interface{}, zerror.Zerror) {
	userId, err := strconv.ParseUint(req.Header.Get("user-id"), 10, 64)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	request := AppUnReleaseRequest{}
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	// check if app is created by user
	app, err := dao.App.GetByAuthorAndId(userId, request.AppId)
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	app.IsPublic = 0
	if err = dao.App.Save(&app); err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}

	return nil, nil
}

func returnError(w http.ResponseWriter, err error) {
	rsp := Response{
		Code:    -1,
		Message: err.Error(),
		Data:    nil,
	}
	json.NewEncoder(w).Encode(rsp)
}

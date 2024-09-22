package logic

import (
	"net/http"
	"strconv"

	"github.com/hankeyyh/chat-box-svr/dao"
)

func AppPublicList(req *http.Request) (interface{}, *zerror) {
	apps, err := dao.App.AllPublic()
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return apps, nil
}

func AppPrivateList(req *http.Request) (interface{}, *zerror) {
	author := req.Form.Get("author")
	if author == "" {
		return nil, NewZError(-1, "author is required", nil)
	}
	// TODO validate string format
	apps, err := dao.App.AllPrivateByAuthor(author)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return apps, nil
}

func AppDetail(req *http.Request) (interface{}, *zerror) {
	appIDStr := req.Form.Get("app_id")
	appId, err := strconv.Atoi(appIDStr)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	app, err := dao.App.GetByID(appId)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	if len(app) == 0 {
		return nil, nil
	}

	return app[0], nil
}

func AppUpsert(req *http.Request) (interface{}, *zerror) {
	return nil, nil
}

func AppRelease(req *http.Request) (interface{}, *zerror) {
	appIDStr := req.Form.Get("app_id")
	appId, err := strconv.Atoi(appIDStr)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	err = dao.App.UpdateIsPublic(appId, true)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return nil, nil
}

func AppUnrelease(req *http.Request) (interface{}, *zerror) {
	appIDStr := req.Form.Get("app_id")
	appId, err := strconv.Atoi(appIDStr)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	err = dao.App.UpdateIsPublic(appId, false)
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return nil, nil
}

func AppChatList(req *http.Request) (interface{}, *zerror) {
	return nil, nil
}

func AppChat(req *http.Request) (interface{}, *zerror) {
	return nil, nil
}
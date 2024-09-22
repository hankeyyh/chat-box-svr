package logic

import (
	"net/http"

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
	err := req.ParseForm()
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
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
	return nil, nil
}

func AppUpsert(req *http.Request) (interface{}, *zerror) {
	return nil, nil
}

func AppRelease(req *http.Request) (interface{}, *zerror) {
	return nil, nil
}

func AppUnrelease(req *http.Request) (interface{}, *zerror) {
	return nil, nil
}

func AppChatList(req *http.Request) (interface{}, *zerror) {
	return nil, nil
}

func AppChat(req *http.Request) (interface{}, *zerror) {
	return nil, nil
}
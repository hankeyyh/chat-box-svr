package logic

import (
	"net/http"

	"github.com/hankeyyh/chat-box-svr/dao"
)

func ModelList(req *http.Request) (interface{}, *zerror) {
	models, err := dao.AiModel.All()
	if err != nil {
		return nil, NewZError(-1, err.Error(), err)
	}
	return models, nil
}

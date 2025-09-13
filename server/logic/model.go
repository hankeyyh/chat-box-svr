package logic

import (
	"net/http"

	"github.com/hankeyyh/chat-box-svr/dao"
	"github.com/hankeyyh/chat-box-svr/zerror"
)

func ModelList(req *http.Request) (interface{}, zerror.Zerror) {
	models, err := dao.AiModel.All()
	if err != nil {
		return nil, zerror.NewZError(-1, err.Error(), err)
	}
	return models, nil
}

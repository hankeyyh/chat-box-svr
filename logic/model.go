package logic

import (
	"encoding/json"
	"net/http"

	"github.com/hankeyyh/chat-box-svr/dao"
)

func ModelList(w http.ResponseWriter, req *http.Request) {
	models, err := dao.AiModel.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(models)
}
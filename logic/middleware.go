package logic

import (
	"encoding/json"
	"net/http"
)

type RequestHandleFunc func(req *http.Request) (interface{}, *zerror)

func HandleResponse(handleFunc RequestHandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := handleFunc(req)
		rsp := Response{
			Code:    err.GetCode(),
			Message: err.GetMessage(),
			Data:    data,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rsp)
	}
}

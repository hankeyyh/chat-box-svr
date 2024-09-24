package logic

import (
	"encoding/json"
	"net/http"
)

type RequestHandleFunc func(req *http.Request) (interface{}, *zerror)

func HandleGetFormRequest(handleFunc RequestHandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Method must be GET"))
			return
		}
		err := req.ParseForm()
		if err != nil {
			rsp := Response{
				Code:    -1,
				Message: err.Error(),
				Data:    nil,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		var zerr *zerror
		data, zerr := handleFunc(req)
		rsp := Response{
			Code:    zerr.GetCode(),
			Message: zerr.GetMessage(),
			Data:    data,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rsp)
	}
}

func HandlePostJsonRequest(handleFunc RequestHandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
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

		var zerr *zerror
		data, zerr := handleFunc(req)
		rsp := Response{
			Code:    zerr.GetCode(),
			Message: zerr.GetMessage(),
			Data:    data,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rsp)
	}
}

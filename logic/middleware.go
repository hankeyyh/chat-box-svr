package logic

import (
	"encoding/json"
	"net/http"
)

type RequestHandleFunc func(req *http.Request) (interface{}, *zerror)

func HandleResponse(handleFunc RequestHandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
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
		}

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

func HandleStreamResponse(handleFunc RequestHandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.(http.Flusher).Flush()
		for {
			data, err := handleFunc(req)
			if data == nil {
				return
			}
			rsp := Response{
				Code:    err.GetCode(),
				Message: err.GetMessage(),
				Data:    data,
			}
			json.NewEncoder(w).Encode(rsp)
			w.(http.Flusher).Flush()
		}
	}
}
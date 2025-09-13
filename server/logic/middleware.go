package logic

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hankeyyh/chat-box-svr/util/log"
	"github.com/hankeyyh/chat-box-svr/zerror"
)

type RequestHandleFunc func(req *http.Request) (interface{}, zerror.Zerror)

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
		log.Infof("request, header: %+v, form: %+v", req.Header, req.Form)
		st := time.Now()

		var zerr zerror.Zerror
		var code int
		var message string
		data, zerr := handleFunc(req)
		if zerr != nil {
			code = zerr.GetCode()
			message = zerr.GetMessage()
		}
		rsp := Response{
			Code:    code,
			Message: message,
			Data:    data,
		}
		elapsed := time.Since(st)
		if code != 0 {
			log.Errorf("response: %+v, duration: %s", rsp, elapsed.String())
		} else {
			log.Infof("response: %+v, duration: %s", rsp, elapsed.String())
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
		log.Infof("request, header: %+v, body: %+v", req.Header, req.Body)
		st := time.Now()

		var zerr zerror.Zerror
		var code int
		var message string
		data, zerr := handleFunc(req)
		if zerr != nil {
			code = zerr.GetCode()
			message = zerr.GetMessage()
		}
		rsp := Response{
			Code:    code,
			Message: message,
			Data:    data,
		}
		elapsed := time.Since(st)
		log.Infof("response: %+v, duration: %s", rsp, elapsed.String())

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rsp)
	}
}

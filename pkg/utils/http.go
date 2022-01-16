package utils

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type HttpHelper struct {
	log *zap.SugaredLogger
}

func NewHttpHelper(log *zap.SugaredLogger) *HttpHelper {
	return &HttpHelper{log: log}
}

func (hh *HttpHelper) WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)

	return err
}

func (hh *HttpHelper) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	data := map[string]interface{}{"error": message}

	if err := hh.WriteJSON(w, status, data); err != nil {
		hh.log.Errorw("server error", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (hh *HttpHelper) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	hh.log.Errorw("server error", "error", err)

	message := "the server encountered a problem and could not process your request"
	hh.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func (hh *HttpHelper) ClientErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	hh.log.Debugw("client error", "error", message)
	hh.ErrorResponse(w, r, status, message)
}

func (hh *HttpHelper) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	hh.ClientErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

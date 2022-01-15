package http

import (
	"encoding/json"
	"net/http"
)

func (fs *FibonacciService) writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/js")
	w.WriteHeader(status)
	_, err = w.Write(js)

	return err
}

func (fs *FibonacciService) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	data := map[string]interface{}{"error": message}

	if err := fs.writeJSON(w, status, data); err != nil {
		fs.log.Errorw("server error", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (fs *FibonacciService) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	fs.log.Errorw("server error", "error", err)

	message := "the server encountered a problem and could not process your request"
	fs.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (fs *FibonacciService) clientErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	fs.log.Debugw("client error", "err", message)
	fs.errorResponse(w, r, status, message)
}

func (fs *FibonacciService) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	fs.clientErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

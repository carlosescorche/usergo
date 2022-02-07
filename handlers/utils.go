package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	errInvalidPayload string = "errInvalidPayload"
)

func Error(w http.ResponseWriter, err ErrorInfo, code int) {
	log.Printf("%v %v", "Has occurred an error: ", err.Message)

	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(code)

	envelope := HTTPResponseError{
		Code:    err.Code,
		Op:      err.Op,
		Message: err.Message,
		Status:  code,
	}

	json.NewEncoder(w).Encode(envelope)
	return
}

func Success(w http.ResponseWriter, message interface{}, code int) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(code)

	envelope := HTTPResponseEnvelope{
		HTTPStatus: code,
		Data:       message,
	}

	json.NewEncoder(w).Encode(envelope)
	return
}

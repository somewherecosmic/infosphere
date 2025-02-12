package utils

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"Msg"`
}

func NewAPIError(statusCode int, msg string) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        msg,
	}
}

func (e *APIError) SendAPIErrorResponse(w http.ResponseWriter) {

	errorMessage := struct {
		Msg string `json:"Msg"`
	}{Msg: e.Msg}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	json.NewEncoder(w).Encode(errorMessage)
}

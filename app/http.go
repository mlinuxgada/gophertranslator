package app

import (
	"net/http"
	"encoding/json"
)

// ErrorResponse - generic err response struct
type ErrorResponse struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

// Render - method for generic json resp rendering
func (App *Application) Render(resp http.ResponseWriter, data interface{}) {

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(resp).Encode(data)
}

// RenderError - method for generic err response render processing
func (App *Application) RenderError(resp http.ResponseWriter, message string, code int) {

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(code)

	err := ErrorResponse{
		Status: code,
		Detail: message,
	}

	json.NewEncoder(resp).Encode(err)
}

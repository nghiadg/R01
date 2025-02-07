package http_utils

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   []Error     `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func WriteJSONResponse(w http.ResponseWriter, response HttpResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}

package util

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

type Response struct {
	Ok    bool   `json:"ok"`
	Data  any    `json:"data"`
	Error string `json:"error"`
}

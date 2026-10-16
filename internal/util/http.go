package util

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	slog.Warn("Path not found.",
		slog.String("path", r.URL.Path),
		slog.String("remote_addr", r.RemoteAddr))

	WriteJson(w, 404, Response{
		Ok:    false,
		Error: errors.New("Path not found.").Error(),
	})
}
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Health checked.", slog.String("remote_addr", r.RemoteAddr))

	WriteJson(w, 200, Response{
		Ok:   true,
		Data: struct{ check_health bool }{check_health: true},
	})
}

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

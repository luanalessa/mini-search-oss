package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Health struct {
	Status   string `json:"status"`
	Service  string `json:"service"`
	Revision string `json:"revision"`
	Env      string `json:"env"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := Health{
		Status:   "ok",
		Service:  "minisearch-api",
		Revision: getenv("GIT_SHA", "dev"),
		Env:      getenv("APP_ENV", "local"),
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

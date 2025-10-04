package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TextHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	HealthHandler(rr, req)

	assert := func(cond bool, msg string, args ...interface{}) {
		if !cond {
			t.Fatalf(msg, args...)
		}
	}

	assert(rr.Code == http.StatusOK, "expected status 200, got %d", rr.Code)

	ct := rr.Header().Get("Content-Type")
	assert(ct == "application/json", "expected application/json, got %s", ct)

	body := rr.Body.String()
	want := `"status":"ok"`
	assert(contains(body, want), "expected body to contain %s, got %s", want, body)

}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > len(sub) && (s[0:len(sub)] == sub || contains(s[1:], sub)))
}

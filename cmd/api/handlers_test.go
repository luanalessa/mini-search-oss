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

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	ct := rr.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Fatalf("expected application/json, got %s", ct)
	}

	body := rr.Body.String()
	if want := `"status":"ok"`; !contains(body, want) {
		t.Fatalf("expected body to contain %s, got %s", want, body)
	}

}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > len(sub) && (s[0:len(sub)] == sub || contains(s[1:], sub)))
}

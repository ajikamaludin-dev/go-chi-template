package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	r := NewRouter(RouterDeps{})
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if rr.Body.String() != "ok" {
		t.Fatalf("expected ok, got %s", rr.Body.String())
	}
}

func TestReadyz(t *testing.T) {
	r := NewRouter(RouterDeps{})
	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}

func TestPing(t *testing.T) {
	r := NewRouter(RouterDeps{})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/ping", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
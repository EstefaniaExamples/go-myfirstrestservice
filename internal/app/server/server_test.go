package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Start(t *testing.T) {

	srv := Server{}
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	srv.ServeHTTP(response, request)

	want := 200
	got := response.Code
	if got != want {
		t.Fatalf("Error starting server, got %d, want %d", got, want)
	}

	wantBody := "ok"
	gotBody := response.Body.String()
	if gotBody != wantBody {
		t.Fatalf("Error starting server, got %q, want %q", gotBody, wantBody)
	}
}
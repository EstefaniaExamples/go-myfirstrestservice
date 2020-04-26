package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {
	srv := NewServer(8080)
	var err error = nil

	go func() {
		err = srv.Start()
	}()

	time.Sleep(100 * time.Millisecond)
	srv.quit()

	time.Sleep(100 * time.Millisecond)
	if err != nil {
		t.Fatalf("Error starting server, got %v, want nil", err)
	}

}

func TestServer_ErrorWhenInvalidPort(t *testing.T)  {
	srv := NewServer(100000000)
	err := srv.Start()

	if err == nil {
		t.Fatalf("expect error got nil")
	}
}

func TestServer_Request(t *testing.T) {
	srv := NewServer(8080)

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()  // mock the response

	srv.ServeHTTP(response, request)

	want := 200
	got := response.Code
	if got != want {
		t.Fatalf("Error requesting server, got %d, want %d", got, want)
	}

	wantBody := "ok"
	gotBody := response.Body.String()
	if gotBody != wantBody {
		t.Fatalf("Error starting server, got %q, want %q", gotBody, wantBody)
	}
}
package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequestReturnsStatusCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	
	client := NewClient()

	statusCode, err := client.Get(server.URL)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if statusCode != http.StatusOK {  
		t.Errorf("Expected status code %d, got %d", http.StatusOK, statusCode)
}
}

package main

import (
	"net/http"
	"testing"
)

func TestStatusCode(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8000/channels/105/programm")

	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

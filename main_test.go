package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandle(t *testing.T) {
	pl, err := ioutil.ReadFile("testdata/payload.json")
	if err != nil {
		t.Fatalf("could not read payload.json: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080", bytes.NewReader(pl))
	if err != nil {
		t.Fatalf("could not create test request: %v", err)
	}

	rec := httptest.NewRecorder()
	handle(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code %d", res.StatusCode)
	}

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("could not read result payload: %v", err)
	}
}

func BenchmarkHandle(b *testing.B) {
	b.StopTimer()

	pl, err := ioutil.ReadFile("testdata/payload.json")
	if err != nil {
		b.Fatalf("could not read payload.json: %v", err)
	}

	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080", bytes.NewReader(pl))
		if err != nil {
			log.Fatalf("could not create test request: %v", err)
		}
		rec := httptest.NewRecorder()
		b.StartTimer()
		handle(rec, req)
		b.StopTimer()
	}
}

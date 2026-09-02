package api

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"struct-ex/config"
	"testing"
)

func TestGetBin_Success(t *testing.T) {
	fakeResponse := `{"record":{"city":"Tokyo"},"metadata":{"id":"abc123"}}`

	originalClient := http.DefaultClient
	http.DefaultClient = &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			if req.Method != "GET" {
				t.Errorf("ожидался метод GET, получен %s", req.Method)
			}
			expectedURL := baseURL + "/abc123/latest"
			if req.URL.String() != expectedURL {
				t.Errorf("ожидался URL %s, получен %s", expectedURL, req.URL.String())
			}
			if req.Header.Get("X-Master-Key") != "test-key" {
				t.Errorf("не установлен X-Master-Key")
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(fakeResponse)),
			}, nil
		}),
	}
	defer func() { http.DefaultClient = originalClient }()

	a := Api{Conf: &config.Config{Key: "test-key"}}

	result, err := a.GetBin("abc123")
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if result != fakeResponse {
		t.Errorf("ожидалось %s, получено %s", fakeResponse, result)
	}
}

func TestGetBin_NetworkError(t *testing.T) {
	originalClient := http.DefaultClient
	http.DefaultClient = &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("connection refused")
		}),
	}
	defer func() { http.DefaultClient = originalClient }()

	a := Api{Conf: &config.Config{Key: "test-key"}}

	_, err := a.GetBin("abc123")
	if err == nil {
		t.Fatal("ожидалась ошибка, но её не было")
	}
}

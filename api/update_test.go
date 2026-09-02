package api

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"struct-ex/config"
	"testing"
)

func TestUpdateBin_Success(t *testing.T) {
	originalClient := http.DefaultClient
	http.DefaultClient = &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			if req.Method != "PUT" {
				t.Errorf("ожидался метод PUT, получен %s", req.Method)
			}
			expectedURL := baseURL + "/abc123"
			if req.URL.String() != expectedURL {
				t.Errorf("ожидался URL %s, получен %s", expectedURL, req.URL.String())
			}
			if req.Header.Get("Content-Type") != "application/json" {
				t.Errorf("не установлен Content-Type")
			}

			sentBody, _ := io.ReadAll(req.Body)
			if string(sentBody) != `{"city":"Osaka"}` {
				t.Errorf("неверное тело запроса: %s", string(sentBody))
			}

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"message":"ok"}`)),
			}, nil
		}),
	}
	defer func() { http.DefaultClient = originalClient }()

	a := Api{Conf: &config.Config{Key: "test-key"}}

	err := a.UpdateBin("abc123", []byte(`{"city":"Osaka"}`))
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
}

func TestUpdateBin_NetworkError(t *testing.T) {
	originalClient := http.DefaultClient
	http.DefaultClient = &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("connection refused")
		}),
	}
	defer func() { http.DefaultClient = originalClient }()

	a := Api{Conf: &config.Config{Key: "test-key"}}

	err := a.UpdateBin("abc123", []byte(`{}`))
	if err == nil {
		t.Fatal("ожидалась ошибка, но её не было")
	}
}

package api

import (
	"io"
	"net/http"
	"strings"
	"struct-ex/config"
	"testing"
)

func TestDeleteBin_Success(t *testing.T) {
	originalClient := http.DefaultClient
	http.DefaultClient = &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			if req.Method != "DELETE" {
				t.Errorf("ожидался метод DELETE, получен %s", req.Method)
			}
			expectedURL := baseURL + "/abc123"
			if req.URL.String() != expectedURL {
				t.Errorf("ожидался URL %s, получен %s", expectedURL, req.URL.String())
			}
			if req.Header.Get("X-Master-Key") != "test-key" {
				t.Errorf("не установлен X-Master-Key")
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"message":"Successful delete"}`)),
			}, nil
		}),
	}
	defer func() { http.DefaultClient = originalClient }()

	a := Api{Conf: &config.Config{Key: "test-key"}}

	err := a.DeleteBin("abc123")
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
}

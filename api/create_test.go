package api

import (
	"io"
	"net/http"
	"strings"
	"struct-ex/config"
	"testing"
)

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestCreateBin_Success(t *testing.T) {
	fakeResponse := `{"metadata": {"id": "abc123", "createdAt": "2024-01-01T00:00:00Z"}}`

	originalClient := http.DefaultClient
	http.DefaultClient = &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			if req.Method != "POST" {
				t.Errorf("ожидался метод POST, получен %s", req.Method)
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

	bin, err := a.CreateBin("my-bin", []byte(`{"city":"Tokyo"}`))
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if bin.ID != "abc123" {
		t.Errorf("ожидался ID abc123, получен %s", bin.ID)
	}
	if bin.Name != "my-bin" {
		t.Errorf("ожидалось имя my-bin, получено %s", bin.Name)
	}
}

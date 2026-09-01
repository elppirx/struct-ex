package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"struct-ex/bins"
	"time"
)

const baseURL = "https://api.jsonbin.io/v3/b"

type createResponse struct {
	Metadata struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
}

func (a Api) CreateBin(name string, body []byte) (*bins.Bin, error) {
	req, err := http.NewRequest("POST", baseURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.Conf.Key)
	req.Header.Set("X-Bin-Name", name)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cr createResponse
	if err := json.Unmarshal(respData, &cr); err != nil {
		return nil, err
	}

	return bins.NewBin(cr.Metadata.ID, false, cr.Metadata.CreatedAt, name), nil
}

func (a Api) GetBin(id string) (string, error) {
	req, err := http.NewRequest("GET", baseURL+"/"+id+"/latest", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Master-Key", a.Conf.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (a Api) UpdateBin(id string, body []byte) error {
	req, err := http.NewRequest("PUT", baseURL+"/"+id, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.Conf.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (a Api) DeleteBin(id string) error {
	req, err := http.NewRequest("DELETE", baseURL+"/"+id, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Master-Key", a.Conf.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

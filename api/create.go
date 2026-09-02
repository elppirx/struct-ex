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

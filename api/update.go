package api

import (
	"bytes"
	"net/http"
)

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

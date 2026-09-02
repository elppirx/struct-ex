package api

import (
	"io"
	"net/http"
)

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

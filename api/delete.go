package api

import "net/http"

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

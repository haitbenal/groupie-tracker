package handlers

import (
	"encoding/json"
	"net/http"
)

func GetJson(URL string, data any) error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(data)
}

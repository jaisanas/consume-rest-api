package util

import (
	"encoding/json"
	"net/http"
	"time"
)

var client *http.Client

func GetJson(url string, target interface{}) error {
	client = &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil  {
		return nil
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
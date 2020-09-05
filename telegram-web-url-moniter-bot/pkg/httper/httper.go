package httper

import (
	"net/http"
)


func GetHTTPInfo(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetHTTPStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func GetHTTPStatus(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "000 FAILED", err
	}
	return resp.Status, nil
}
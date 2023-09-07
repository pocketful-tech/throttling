package helpers

import (
	"bytes"
	"net/http"
)

func CallApi(methodType, url string, payload *bytes.Buffer, authToken string) (*http.Response, error) {
	var req *http.Request
	if payload == nil {
		req, _ = http.NewRequest(methodType, url, nil)
	} else {
		req, _ = http.NewRequest(methodType, url, payload)
	}
	req.Header.Set("Authorization", authToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}

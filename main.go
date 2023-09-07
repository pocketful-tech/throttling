package main

import (
	"bytes"
	"io"
	"throttling/helpers"
	"time"
)

func MakeAPIRequest(method string, apiURL string, payload *bytes.Buffer, authorization string) ([]byte, error) {

	var resByte []byte

	resp, err := helpers.CallApi(method, apiURL, payload, authorization)

	if err != nil {
		return resByte, err
	}
	defer resp.Body.Close()

	// Read and print the response body (you can handle the response as needed)
	resByte, err = io.ReadAll(resp.Body)
	if err != nil {
		return resByte, err
	}

	return resByte, nil
}

func NewAPIThrottler(requestsPerSecond int, clientName, clientIp, vendorName string) *helpers.APIThrottler {
	return &helpers.APIThrottler{
		RequestsPerSecond: requestsPerSecond,
		LastRequestTime:   time.Now(),
		ClientName:        clientName,
		ClientIp:          clientIp,
		VendorName:        vendorName,
	}
}

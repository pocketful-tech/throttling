package throttling

import (
	"bytes"
	"io"
	"time"
)

func MakeAPIRequest(method string, apiURL string, payload *bytes.Buffer, authorization string) ([]byte, error) {

	var resByte []byte

	// Create an HTTP GET request

	// req, err := http.NewRequest(method, apiURL, payload)
	// if err != nil {
	// 	return resByte, err
	// }

	// // Send the request
	// client := &http.Client{}
	// resp, err := client.Do(req)

	resp, err := CallApi(method, apiURL, payload, authorization)

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

func NewAPIThrottler(requestsPerSecond int, clientName, clientIp, vendorName string) *APIThrottler {
	return &APIThrottler{
		RequestsPerSecond: requestsPerSecond,
		LastRequestTime:   time.Now(),
		ClientName:        clientName,
		ClientIp:          clientIp,
		VendorName:        vendorName,
	}
}

package throttling

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func MakeAPIRequest(apiURL string) ([]byte, error) {

	var resByte []byte

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error1: ", err)
		return resByte, err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
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

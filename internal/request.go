package internal

import (
	"net/http"
)

// Request to url
func RequestReceiveing(url, token string) (*http.Response, error) {

	Request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Sending token and receiving data
	Request.Header.Set("Client-Id", "627")
	Request.Header.Set("Client-Secret", token)
	client := &http.Client{}
	Response, err := client.Do(Request)
	if err != nil {
		return nil, err
	}

	return Response, nil // return for encodingJson
}

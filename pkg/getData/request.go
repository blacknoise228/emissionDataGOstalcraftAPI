package getData

import (
	"net/http"
	"stalcraftbot/internal/logs"
)

// Request to StalcraftAPI and returned json response
func RequestReceiveing(url, clientID, token string) (*http.Response, error) {

	Request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Sending token and receiving data
	Request.Header.Set("Client-Id", clientID)
	Request.Header.Set("Client-Secret", token)
	client := &http.Client{}
	Response, err := client.Do(Request)
	if err != nil {
		return nil, err
	}
	logs.Logger.Debug().Msg("RequestReceiveing done")
	return Response, nil // return for encodingJson
}

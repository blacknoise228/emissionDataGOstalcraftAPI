package jsWorker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Decoding json to structure
func EncodingJson(resp *http.Response) (EmissionInfo, error) {
	defer resp.Body.Close()
	var EmissionData EmissionInfo
	if err := json.NewDecoder(resp.Body).Decode(&EmissionData); err != nil {
		fmt.Println(err)
		return EmissionData, err
	}
	return EmissionData, nil
}

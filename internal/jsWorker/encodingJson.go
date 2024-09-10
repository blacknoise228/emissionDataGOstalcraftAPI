package jsWorker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structure .json
type EmissionInfo struct {
	CurrentStart  string `json:"currentStart"`  // last emission time
	PreviousStart string `json:"previousStart"` // preview emission time
	PreviousEnd   string `json:"previousEnd"`   // preview emission end
	Status        int    `json:"status"`        // status normal = 0, if status = 401, recreate token auth
}

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

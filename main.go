package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Structure .json
type EmissionInfo struct {
	CurrentStart  string `json:"currentStart"`  // last emission time
	PreviousStart string `json:"previousStart"` // preview emission time
	PreviousEnd   string `json:"previousEnd"`   // preview emission end
	Status        int    `json:"status"`        // status normal = 0, if status = 401, recreate token auth
}

func main() {
	// this case show you work with demoAPI. you have to change to the actual token and url

	url := "https://dapi.stalcraft.net/ru/emission"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwic3ViIjoiMSIsIm5iZiI6MTY3Mzc5NzgzOCwiZXhwIjo0ODI3Mzk3ODM4LCJpYXQiOjE2NzM3OTc4MzgsImp0aSI6IjJlamRwOG54a3A1djRnZWdhbWVyeWlkMW5ic24zZDhpZ2oyejgzem1vMDYzNjNoaXFkNWhwOTY1MHZwdWh4OXEybXBmd2hnbnUxNHR5cmp2In0.Ocw4CzkkuenkAOjkAR1RuFgLqix7VJ-8vWVS3KAJ1T3SgIWJG145xqG2qms99knu5azn_oaoeyMOXhyG_fuMQFGOju317GiS6pAXAFGOKvxcUCfdpFcEHO6TWGM8191-tlfV-0rAqCi62gprKyr-SrUG3nUJhv6XKegja_vYVujRVx0ouAaDvDKawiOssG5If_hXGhdhnmb3_7onnIc4hFsm4i9QVkWXe8GO6OsS999ZIX0ClNhTk2kKKTl2dDVIiKha_HB1aghm_LOYoRgb3i3B_DH4UO312rHYR5I4qO43c8x-TW7NwovItDSzhiCmcxZuUUeAUF3yFr5ovaR4fMj1LEy3y3V2piQDKPwmBOpI9S6OzWUIBJYcRYlT2HIrWCRc0YvM7AOGoxcH2Gf4ncqcF_M8fw7IMKf3pdnuxf1EbdEpzOapBD1Pw065em-U8PN4LVzw9lhIHx_Yj69qaFEx7Bhw3BCwsrx-o9hgg7T1TOV6kF11YfR99lIuj9z96XBLg5ipt-M_j7nHRoHWhM0Rc6uLIKPg0In0xYkybSfWG6v3Hs6kwgB7wkqpXpoVQltJvlqjtlf9Pp4zmkqlWQHx9as4xsgoTAQyCgaC0kisICNC58_g3QrJAfoFXW68x-OHlRKCAPqoR9V-0cVs-B83szaFmsEGegAttFLlDhE"

	// request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	// Sending token and receiving data
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
	}
	// Decoding json to structure
	var emissionData EmissionInfo
	if err := json.NewDecoder(resp.Body).Decode(&emissionData); err != nil {
		fmt.Println(err)
	}

	// Last emission time start
	lastEmissionStart, err := time.Parse(time.RFC3339Nano, emissionData.PreviousStart)
	if err != nil {
		fmt.Println(err)
	}
	lastEmissionStart = lastEmissionStart.In(time.Local) //convert to your time zone

	// Last emission time end
	lastEmissionEnd, err := time.Parse(time.RFC3339Nano, emissionData.PreviousEnd)
	if err != nil {
		fmt.Println(err)
	}
	lastEmissionEnd = lastEmissionEnd.In(time.Local) //convert to your time zone

	// Time after last emission start
	timeDurNow := time.Since(lastEmissionStart).Round(time.Second)

	// Print Result
	fmt.Printf("Last emission start/end: %v - %v\nTime after last emission: %v\n", lastEmissionStart.Format(time.DateTime), lastEmissionEnd.Format(time.DateTime), timeDurNow)

}

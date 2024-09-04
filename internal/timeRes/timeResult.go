package timeRes

import (
	"fmt"
	"time"

	jSon "stalcraftBot/internal/jSon"
)

// Work with time data output for user
func TimeResult(data jSon.EmissionInfo) (string, error) {

	// Last emission time start
	lastEmissionStart, err := time.Parse(time.RFC3339Nano, data.PreviousStart)
	if err != nil {
		return "", err
	}
	lastEmissionStart = lastEmissionStart.In(time.Local) //convert to your time zone

	// Last emission time end
	lastEmissionEnd, err := time.Parse(time.RFC3339Nano, data.PreviousEnd)
	if err != nil {
		return "", err
	}
	lastEmissionEnd = lastEmissionEnd.In(time.Local) //convert to your time zone

	// Time after last emission start
	timeDurNow := time.Since(lastEmissionStart).Round(time.Second)

	// Print Result
	return fmt.Sprintf(
		"\nНачало последнего выброса: \n%v\nКонец последнего выброса: \n%v\nПрошло времени с окончания последнего выброса: \n%v\n",
		lastEmissionStart.Format(time.DateTime),
		lastEmissionEnd.Format(time.DateTime),
		timeDurNow,
	), nil
}
func CurrentEmissionResult(data jSon.EmissionInfo) (string, error) {
	currentEmissionStart, err := time.Parse(time.RFC3339Nano, data.CurrentStart)
	if err != nil {
		return "", err
	}
	currentEmissionStart = currentEmissionStart.In(time.Local)
	return fmt.Sprintf("\nВсем кто меня слышит! Приближается выброс! Срочно ищите себе укрытие!\n%v", currentEmissionStart.Format(time.DateTime)), nil
}

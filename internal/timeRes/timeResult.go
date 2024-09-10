package timeRes

import (
	"fmt"
	"stalcraftBot/internal/jsWorker"
	"stalcraftBot/internal/logs"

	"time"
)

// Processing last emission data json to a given date and time format, with comments about emission for users
func TimeResult(data jsWorker.EmissionInfo) (string, error) {

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
	logs.Logger.Debug().Msg("TimeResult done")

	// Print Result
	return fmt.Sprintf(
		"\nНачало последнего выброса: \n%v\nКонец последнего выброса: \n%v\nПрошло времени с окончания последнего выброса: \n%v\n",
		lastEmissionStart.Format(time.DateTime),
		lastEmissionEnd.Format(time.DateTime),
		timeDurNow,
	), nil
}

// Processing current emission data json to a given date and time format, with comments about emission for users
func CurrentEmissionResult(data jsWorker.EmissionInfo) (string, error) {
	currentEmissionStart, err := time.Parse(time.RFC3339Nano, data.CurrentStart)
	if err != nil {
		return "", err
	}
	currentEmissionStart = currentEmissionStart.In(time.Local)
	logs.Logger.Debug().Msg("CurrentEmissionResult done")
	return fmt.Sprintf("\nВсем кто меня слышит! Приближается выброс! Срочно ищите себе укрытие!\n%v", currentEmissionStart.Format(time.DateTime)), nil
}

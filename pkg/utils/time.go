package utils

import (
	"time"
)

func GetTimezoneOffset(timezone string) (time.Duration, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}
	// Get the current time in the specified timezone
	currentTime := time.Now().In(loc)
	// Get the timezone offset in seconds
	return currentTime.UTC().Sub(currentTime), nil
}

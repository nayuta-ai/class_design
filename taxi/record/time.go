package record

import (
	"strconv"
	"strings"
)

type Time struct {
	Hours   int64
	Minutes float64
}

// newTime converts the timestamp into Time class.
// time.Hours is hours and time.Minutes is minutes.
func newTime(timestamp string) (Time, error) {
	timeObject := strings.Split(timestamp, ":")
	var time Time
	hours, err := extractHours(timeObject)
	if err != nil {
		return Time{}, err
	}
	time.Hours = hours
	minutes, err := ExtractMinutes(timeObject)
	if err != nil {
		return Time{}, err
	}
	time.Minutes = minutes
	return time, nil
}

// extractHours extracts the timeObject into hours
// whose type is int64
func extractHours(timeObject []string) (int64, error) {
	hoursString := timeObject[0]
	hoursInt64, err := strconv.ParseInt(hoursString, 10, 64)
	if err != nil {
		return 0, err
	}
	return hoursInt64, nil
}

// extractMinutes extracts the timeObject into minutes
// whose type is float64
func ExtractMinutes(timeObject []string) (float64, error) {
	var (
		totalMinutes float64 = 0
		timeConverter float64 = 3600
	)
	const timeUnit float64 = 60
	for _, timeString := range timeObject {
		timeFloat, err := strconv.ParseFloat(timeString, 64)
		if err != nil {
			return 0, err
		}
		totalMinutes += timeFloat * timeConverter
		timeConverter /= timeUnit
	}
	return totalMinutes, nil
}

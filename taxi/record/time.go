package record

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hours   int64
	Minutes float64
}

const (
	lengthTimestamp int   = 3
	maxHour         int64 = 100
	minHour         int64 = 0
)

// newTime creates a Time instance from timestamp string.
// time.Hours is hours and time.Minutes is minutes.
func newTime(timestamp string) (Time, error) {
	timeObject := strings.Split(timestamp, ":")
	if len(timeObject) != lengthTimestamp {
		return Time{}, errors.New("the length of a list of timeObject should be 3")
	}
	var time Time
	hours, err := extractHours(timeObject)
	if err != nil {
		return Time{}, err
	}
	if hours < minHour || maxHour <= hours {
		return Time{}, fmt.Errorf("hours should be more than equal to %d and less than %d", minHour, maxHour)
	}
	time.Hours = hours
	minutes, err := ExtractMinutes(timeObject)
	if err != nil {
		return Time{}, err
	}
	time.Minutes = minutes
	return time, nil
}

// extractMinutes extracts minutes info from timeObject
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
		totalMinutes  float64 = 0
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

package parser

import (
	"strconv"
	"strings"
	"taxi/record"
)

// Parse can parse the taxi logs. We usually used it  
// when caclurating some information from taxi logs.
func Parse(logs string) ([]record.Record, error) {
	var records []record.Record
	for _, log := range strings.Split(logs, "<LF>") {
		// Skip a blank line
		if log == ""{
			break
		}
		convertedLog := strings.Split(log, " ")
		var rec record.Record
		time, err := createTimeObject(convertedLog[0])
		if err != nil {
			return []record.Record{}, err
		}
		rec.Time = time
		distance, err := strconv.ParseFloat(convertedLog[1], 64)
		if err != nil {
			return []record.Record{}, err
		}
		rec.Distance = distance
		records = append(records, rec)
	}
	return records, nil
}

// createTimeObject converts the timestamp into Time class.
// time.Hours is hours and time.Minutes is minutes.
func createTimeObject(timestamp string) (record.Time, error) {
	timeObject := strings.Split(timestamp, ":")
	var time record.Time
	hours, err := convertTimeIntoHours(timeObject)
	if err != nil {
		return record.Time{}, err
	}
	time.Hours = hours
	minutes, err := convertTimeIntoMinutes(timeObject)
	if err != nil {
		return record.Time{}, err
	}
	time.Minutes = minutes
	return time, nil
}

// convertTimeIntoHours converts the timeObject into hours
// whose type is int64
func convertTimeIntoHours(timeObject []string) (int64, error) {
	hoursString := timeObject[0]
	hoursInt64, err := strconv.ParseInt(hoursString, 10, 64)
	if err != nil {
		return 0, err
	}
	return hoursInt64, nil
}

// convertTimeIntoMinutes converts the timeObject into minutes
// whose type is float64
func convertTimeIntoMinutes(timeObject []string) (float64, error) {
	var totalMinutes float64
	var timeConverter float64 = 3600
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

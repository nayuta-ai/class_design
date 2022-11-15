package record

import (
	"strconv"
	"strings"
)

type Record struct {
	Distance float64
	Time     Time
}

// NewRecords can create the new record from the taxi logs.
// We usually used it when caclurating some information from taxi logs.
func NewRecords(logs string) ([]Record, error) {
	var records []Record
	for _, log := range strings.Split(logs, "<LF>") {
		// Skip a blank line
		if log == "" {
			break
		}
		splittedLog := strings.Split(log, " ")
		var rec Record
		time, err := newTime(splittedLog[0])
		if err != nil {
			return []Record{}, err
		}
		rec.Time = time
		distance, err := strconv.ParseFloat(splittedLog[1], 64)
		if err != nil {
			return []Record{}, err
		}
		rec.Distance = distance
		records = append(records, rec)
	}
	return records, nil
}

package record

import (
	"errors"
	"strconv"
	"strings"
)

type Record struct {
	Distance float64
	Time     Time
}

const lengthSplittedLog int = 2

// NewRecords can create a list of Record instances from the taxi logs.
// We usually used it when caclurating some information from taxi logs.
func NewRecords(logs string) ([]Record, error) {
	var records []Record
	for _, log := range strings.Split(logs, "<LF>") {
		// Skip a blank line
		if log == "" {
			break
		}
		splittedLog := strings.Split(log, " ")
		if len(splittedLog) != lengthSplittedLog {
			return []Record{}, errors.New("the length of a list of splittedLog should be 2")
		}
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

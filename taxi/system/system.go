package system

import (
	"strconv"
	"taxi/record"
)

// CalcTotalFare calculates a total fare based on logs.
// Also, it is the entry point function of the taxi system.
// If a taxi drives midnight, caused additional fare.
// If a taxi drives under 10 km/h, caused additional fare.
func CalcTotalFare(logs string) (int64, error) {
	records, err := record.NewRecords(logs)
	if err != nil {
		return 0, err
	}
	totalDistance, midnightMinutes := fetchInfoFromLog(records)
	totalFare := calcBaseFare(totalDistance) + calcAdditionalFare(midnightMinutes)
	return totalFare, nil
}

// fetchInfoFromLog fetches the total distance and sum of the midnight minutes.
func fetchInfoFromLog(records []record.Record) (float64, float64) {
	var midnightMinutes float64
	var totalDistance float64
	for i := 1; i < len(records); i++ {
		prevRecord := records[i-1]
		currRecord := records[i]
		currMidnightMinutes := fetchMidnightTime(prevRecord.Time, currRecord.Time)
		if calcVelocity(currRecord.Time.Minutes-prevRecord.Time.Minutes, currRecord.Distance) <= float64(10) {
			totalDistance += currRecord.Distance * 1.25
			midnightMinutes += currMidnightMinutes * 1.25
		} else {
			totalDistance += currRecord.Distance
			midnightMinutes += currMidnightMinutes
		}
	}
	return totalDistance, midnightMinutes
}

// calcAdditionalFare calculates a fare based on minutes during midnight.
func calcAdditionalFare(midnightMinutes float64) int64 {
	const (
		timeUnit int64 = 90
		payUnit  int64 = 80
	)
	return int64(midnightMinutes) / timeUnit * payUnit
}

// calcBaseFare calculates a fare based on distance.
func calcBaseFare(distance float64) int64 {
	var baseFare int = 410
	const (
		distanceInit int = 1053
		distanceUnit int = 237
		payUnit      int = 80
	)
	convertedDistance := int(distance)
	if convertedDistance < distanceInit {
		return int64(baseFare)
	} else {
		convertedDistance -= distanceInit
		baseFare += (convertedDistance/distanceUnit + 1) * payUnit
	}
	return int64(baseFare)
}

// fetchMidnightTime fetches the time driven during midnight.
func fetchMidnightTime(prevTime record.Time, currTime record.Time) float64 {
	midnightTimes := [][]int64{{0, 5}, {22, 29}, {46, 53}, {70, 77}, {94, 99}}
	var totalMidnightMinutes float64
	var (
		ptnHour    int64   = prevTime.Hours
		ptnMinutes float64 = prevTime.Minutes
	)
	for _, midnightTime := range midnightTimes {
		if currTime.Hours <= midnightTime[0] {
			break
		} else if currTime.Hours <= midnightTime[1] && ptnHour < midnightTime[0] {
			midpointMinutes, _ := record.ExtractMinutes([]string{strconv.FormatInt(midnightTime[0], 10), "00", "00.000"})
			totalMidnightMinutes += currTime.Minutes - midpointMinutes
			return totalMidnightMinutes
		} else if currTime.Hours <= midnightTime[1] && midnightTime[0] <= ptnHour {
			totalMidnightMinutes += currTime.Minutes - ptnMinutes
			return totalMidnightMinutes
		} else if midnightTime[1] < currTime.Hours && ptnHour < midnightTime[0] {
			headMidpointMinutes, _ := record.ExtractMinutes([]string{strconv.FormatInt(midnightTime[0], 10), "00", "00.000"})
			tailMidpointMinutes, _ := record.ExtractMinutes([]string{strconv.FormatInt(midnightTime[1], 10), "00", "00.000"})
			totalMidnightMinutes += tailMidpointMinutes - headMidpointMinutes
			ptnMinutes = tailMidpointMinutes
			ptnHour = midnightTime[1]
		} else if midnightTime[1] < currTime.Hours && ptnHour < midnightTime[1] {
			midpointMinutes, _ := record.ExtractMinutes([]string{strconv.FormatInt(midnightTime[1], 10), "00", "00.000"})
			totalMidnightMinutes += midpointMinutes - ptnMinutes
			ptnMinutes = midpointMinutes
			ptnHour = midnightTime[1]
		}
	}
	return totalMidnightMinutes
}

// calcVelocity calculates the velocity of each line in the log.
// We usually use it when checking whether a taxi drives at a low speed.
func calcVelocity(diffMinutes float64, distance float64) float64 {
	const (
		convertedMIntoKm          float64 = 1000
		convertedMinutesIntoHours float64 = 3600
	)
	kmDistance := distance / convertedMIntoKm
	diffHours := diffMinutes / convertedMinutesIntoHours
	return kmDistance / diffHours
}

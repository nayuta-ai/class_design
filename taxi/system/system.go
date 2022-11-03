package system

import (
	"taxi/parser"
)

// CalcTotalFare calcurates a total fare based on logs. 
// If a taxi drives midnight, caused additional fare.
// If a taxi drives under 10 km/h, caused additional fare.
func CalcTotalFare(logs string) (int64, error) {
	var midnightMinutes float64
	var totalDistance float64
	var totalFare int64
	records, err := parser.Parse(logs)
	if err != nil {
		return 0, err
	}
	for i:=1; i < len(records); i++ {
		prevMinutes := records[i-1].Time.Minutes
		currMinutes := records[i].Time.Minutes
		var currMidnightMinutes float64 = 0
		if checkDrivingDuringMidnight(
			records[i-1].Time.Hours,
			records[i].Time.Hours,
		){
			currMidnightMinutes += currMinutes - prevMinutes
		}
		if calcVelocity(currMinutes - prevMinutes, records[i].Distance) <= float64(10) {
			totalDistance += records[i].Distance * 1.25
			midnightMinutes += currMidnightMinutes * 1.25
		} else {
			totalDistance += records[i].Distance
			midnightMinutes += currMidnightMinutes
		}
	}
	totalFare += calcAdditionalFare(midnightMinutes)
	totalFare += calcBaseFare(totalDistance)
	return totalFare, nil
}

// calcAdditionalFare calcurates a fare based on minutes during midnight.
func calcAdditionalFare(midnightMinutes float64) int64 {
	const timeUnit, payUnit int64 = 90, 80
	return int64(midnightMinutes) / timeUnit * payUnit
}

// calcBaseFare calcurates a fare based on distance.
func calcBaseFare(distance float64) int64 {
	var pay int = 410
	const distanceInit, distanceUnit, payUnit int = 1053, 237, 80
	convertedDistance := int(distance)
	if convertedDistance < distanceInit {
		return int64(pay)
	}else{
		convertedDistance -= distanceInit
		pay += (convertedDistance/distanceUnit+1) * payUnit
	}
	return int64(pay)
}

// checkDrivingDuringMidnight check whether the driving during midnight exists.
func checkDrivingDuringMidnight(prevHours int64, currHours int64) bool {
	midnightTimes := [][]int64{{0, 5}, {22, 29}}
	for _, midnightTime := range midnightTimes {
		if midnightTime[0] <= prevHours && prevHours < midnightTime[1] && midnightTime[0] <= currHours && currHours < midnightTime[1]{
			return true
		}
	}
	return false
}

// calcVelocity calcurates velocity each line in the log.
// We usually use it when checking whether a taxi drives low speed.
func calcVelocity(diffMinutes float64, distance float64) float64 {
	const convertedMIntoKm, convertedMinutesIntoHours float64 = 1000, 3600
	kmDistance := distance / convertedMIntoKm
	diffHours := diffMinutes / convertedMinutesIntoHours
	return kmDistance / diffHours
}

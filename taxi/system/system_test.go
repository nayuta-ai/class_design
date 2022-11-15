package system

import (
	"taxi/record"
	"testing"
)

func TestCalcTotalFare(t *testing.T) {
	tests := []struct {
		name string
		test string
		want int64
	}{
		{
			name: "test1",
			test: "13:50:08.245 0.0<LF>13:50:11.123 4.0<LF>13:50:12.125 10.2<LF>13:50:13.100 8.7<LF>",
			want: 410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalFare, err := CalcTotalFare(tt.test)
			if err != nil {
				t.Error(err)
			}
			if totalFare != tt.want {
				t.Errorf("The totalFare should be %v.", tt.want)
			}
		})
	}
}

func TestFetchInfoFromLog(t *testing.T) {
	tests := []struct {
		name string
		test []record.Record
		want []float64
	}{
		{
			name: "test1",
			test: []record.Record{
				{
					Distance: 0,
					Time: record.Time{
						Minutes: 49808.245,
						Hours:   13,
					},
				},
				{
					Distance: 4,
					Time: record.Time{
						Minutes: 49811.123,
						Hours:   13,
					},
				},
				{
					Distance: 10.2,
					Time: record.Time{
						Minutes: 49812.125,
						Hours:   13,
					},
				},
				{
					Distance: 8.7,
					Time: record.Time{
						Minutes: 49813.100,
						Hours:   13,
					},
				},
			},
			want: []float64{
				23.9,
				0.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalDistance, midnightMinutes := fetchInfoFromLog(tt.test)
			if totalDistance != tt.want[0] {
				t.Errorf("The total distance should be %v, but %v", tt.want[0], totalDistance)
			}
			if midnightMinutes != tt.want[1] {
				t.Errorf("The total midnight times should be %v, but %v", tt.want[1], midnightMinutes)
			}
		})
	}
}

func TestCalcMidnightMinutes(t *testing.T) {
	tests := []struct {
		name string
		test float64
		want int64
	}{
		{
			name: "test1",
			test: 89,
			want: 0,
		},
		{
			name: "test2",
			test: 90,
			want: 80,
		},
		{
			name: "test3",
			test: 180,
			want: 160,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fare := calcAdditionalFare(tt.test)
			if fare != tt.want {
				t.Errorf("The fare should be %v.", tt.want)
			}
		})
	}
}

func TestCalcBaseFare(t *testing.T) {
	tests := []struct {
		name string
		test float64
		want int64
	}{
		{
			name: "test1",
			test: 1052,
			want: 410,
		},
		{
			name: "test2",
			test: 1053,
			want: 490,
		},
		{
			name: "test3",
			test: 1290,
			want: 570,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fare := calcBaseFare(tt.test)
			if fare != tt.want {
				t.Errorf("The fare should be %v.", tt.want)
			}
		})
	}
}

func TestFetchMidnightTime(t *testing.T) {
	tests := []struct {
		name string
		test []record.Time
		want float64
	}{
		{
			name: "test1",
			test: []record.Time{
				{Minutes: 0.0, Hours: 0},
				{Minutes: 21600, Hours: 6},
			},
			want: 18000,
		},
		{
			name: "test2",
			test: []record.Time{
				{Minutes: 0.0, Hours: 0},
				{Minutes: 82800, Hours: 23},
			},
			want: 21600,
		},
		{
			name: "test3",
			test: []record.Time{
				{Minutes: 36000.0, Hours: 10},
				{Minutes: 108000.0, Hours: 30},
			},
			want: 25200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currMidnightMinutes := fetchMidnightTime(tt.test[0], tt.test[1])
			if currMidnightMinutes != tt.want {
				t.Errorf("The midnight minutes should be %v, but %v", tt.want, currMidnightMinutes)
			}
		})
	}
}

func TestCalcVelocity(t *testing.T) {
	type test struct {
		diffMinutes float64
		distance    float64
	}
	tests := []struct {
		name string
		test test
		want float64
	}{
		{
			name: "test1",
			test: test{
				diffMinutes: 2.878,
				distance:    4,
			},
			want: 5.0034746351633075,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			velocity := calcVelocity(tt.test.diffMinutes, tt.test.distance)
			if velocity != tt.want {
				t.Errorf("The velocity should be %v.", tt.want)
			}
		})
	}
}

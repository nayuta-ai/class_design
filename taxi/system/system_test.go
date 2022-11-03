package system

import (
	"testing"
)

func TestCalcTotalFare(t *testing.T){
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
		t.Run(tt.name, func(t *testing.T){
			totalFare, err := CalcTotalFare(tt.test)
			if err != nil {
				t.Error(err)
			}
			if totalFare != tt.want{
				t.Errorf("The totalFare should be %v.", tt.want)
			}
		})
	}
}

func TestCalcMidnightMinutes(t *testing.T){
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
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			fare := calcAdditionalFare(tt.test)
			if fare != tt.want{
				t.Errorf("The fare should be %v.", tt.want)
			}
		})
	}
}

func TestCalcBaseFare(t *testing.T){
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
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			fare := calcBaseFare(tt.test)
			if fare != tt.want{
				t.Errorf("The fare should be %v.", tt.want)
			}
		})
	}
}

func TestCheckDrivingDuringMidnight(t *testing.T){
	type test struct {
		prevHours int64
		currHours int64
	}
	tests := []struct {
		name string
		test test
		want bool
	}{
		{
			name: "test1",
			test: test{
				prevHours: 13,
				currHours: 13,
			},
			want: false,
		},
	}
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			ok := checkDrivingDuringMidnight(tt.test.prevHours, tt.test.currHours)
			if ok != tt.want{
				t.Errorf("The function should return %v.", tt.want)
			}
		})
	}
}

func TestCalcVelocity(t *testing.T){
	type test struct {
		diffMinutes float64
		distance float64
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
				distance: 4,
			},
			want: 5.0034746351633075,
		},
	}
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			velocity := calcVelocity(tt.test.diffMinutes, tt.test.distance)
			if velocity != tt.want{
				t.Errorf("The velocity should be %v.", tt.want)
			}
		})
	}
}
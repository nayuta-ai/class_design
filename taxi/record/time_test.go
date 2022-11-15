package record

import (
	"testing"
)

func TestCreateTimeObject(t *testing.T) {
	tests := []struct {
		name string
		test string
		want Time
	}{
		{"test1", "13:50:08.245", Time{
			Hours:   13,
			Minutes: 49808.245,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeObject, err := newTime(tt.test)
			if err != nil {
				t.Error(err)
			}
			if timeObject.Hours != tt.want.Hours {
				t.Errorf("The timeObject.Hours should be %v.", tt.want.Hours)
			}
			if timeObject.Minutes != tt.want.Minutes {
				t.Errorf("The timeObject.Minutes should be %v.", tt.want.Minutes)
			}
		})
	}
}

func TestConvertTimeIntoHours(t *testing.T) {
	tests := []struct {
		name string
		test []string
		want int64
	}{
		{"test1", []string{"13", "50", "08.245"}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hours, err := extractHours(tt.test)
			if err != nil {
				t.Error(err)
			}
			if hours != tt.want {
				t.Errorf("The minutes should be %v.", tt.want)
			}
		})
	}
}

func TestConvertTimeIntoMinutes(t *testing.T) {
	tests := []struct {
		name string
		test []string
		want float64
	}{
		{"test1", []string{"13", "50", "08.245"}, 49808.245},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minutes, err := ExtractMinutes(tt.test)
			if err != nil {
				t.Error(err)
			}
			if minutes != tt.want {
				t.Errorf("The minutes should be %v.", tt.want)
			}
		})
	}
}

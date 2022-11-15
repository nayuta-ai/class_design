package record

import (
	"fmt"
	"testing"
)

func TestNewRecord(t *testing.T) {
	tests := []struct {
		name string
		test string
		want []Record
	}{
		{
			name: "test1",
			test: "13:50:08.245 0.0",
			want: []Record{
				{
					Distance: 0.0,
					Time: Time{
						Minutes: 49808.245,
						Hours:   13,
					},
				},
			},
		},
		{
			name: "test2",
			test: "13:50:08.245 0.0<LF>13:50:11.123 4.0<LF>13:50:12.125 10.2<LF>13:50:13.100 8.7<LF>",
			want: []Record{
				{
					Distance: 0,
					Time: Time{
						Minutes: 49808.245,
						Hours:   13,
					},
				},
				{
					Distance: 4,
					Time: Time{
						Minutes: 49811.123,
						Hours:   13,
					},
				},
				{
					Distance: 10.2,
					Time: Time{
						Minutes: 49812.125,
						Hours:   13,
					},
				},
				{
					Distance: 8.7,
					Time: Time{
						Minutes: 49813.100,
						Hours:   13,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			records, err := NewRecords(tt.test)
			fmt.Println(records)
			if err != nil {
				t.Error(err)
			}
			for i := 0; i < len(records); i++ {
				fmt.Println(i)
				if records[i].Distance != tt.want[i].Distance {
					t.Errorf("The records[i].Distance should be %v.", tt.want[i].Distance)
				}
				if records[i].Time.Hours != tt.want[i].Time.Hours {
					t.Errorf("The records[i].Time.Hours should be %v.", tt.want[i].Time.Hours)
				}
				if records[i].Time.Minutes != tt.want[i].Time.Minutes {
					t.Errorf("The records[i].Time.Minutes should be %v.", tt.want[i].Time.Minutes)
				}
			}
		})
	}
}

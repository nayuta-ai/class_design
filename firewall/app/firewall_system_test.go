package app

import (
	"firewall/request"
	"fmt"
	"testing"
)

func TestAddService(t *testing.T) {
	tests := []struct {
		name string
		test string
	}{
		{"test1", "dns"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddService(tt.test)
			if err != nil {
				t.Fatal(err)
			}
			_, ok := NameService[tt.test]
			if !ok {
				t.Fatalf("Not add service")
			}
		})
	}
}

func TestDeleteService(t *testing.T) {
	tests := []struct {
		name string
		test string
	}{
		{"test1", "dns"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddService(tt.test)
			if err != nil {
				t.Fatal(err)
			}
			deleteService(tt.test)
			_, ok := NameService[tt.test]
			if ok {
				t.Fatalf("Not delete service")
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type test struct {
		service_name           string
		protocol               string
		source_ip_address      string
		destination_ip_address string
		source_port            string
		destination_port       string
	}
	tests := []struct {
		name string
		test test
	}{
		{
			name: "test1",
			test: test{"dns", "tcp", "XXX.XXX.XXX.XXX", "XXX.XXX.XXX.XXX", "80", "53"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.test
			requestJSON := fmt.Sprintf(
				`{
					"service_name":"%v",
					"protocol":"%v",
					"source_ip_address":"%v",
					"destination_ip_address":"%v",
					"source_port":"%v",
					"destination_port":"%v"
				}`,
				want.service_name,
				want.protocol,
				want.source_ip_address,
				want.destination_ip_address,
				want.source_port,
				want.destination_port)
			req, err := request.ParseRequest(requestJSON)
			if err != nil {
				t.Fatal(err)
			}
			if filter(req) {
				t.Fatalf("Unexpected Action: Not pass request")
			}
			err = AddService(want.service_name)
			if err != nil {
				t.Fatal(err)
			}

			if !filter(req) {
				t.Fatalf("Unexpected Action: Should pass request")
			}
			deleteService(want.service_name)
			if filter(req) {
				t.Fatalf("Unexpected Action: Not pass request")
			}
		})
	}
}

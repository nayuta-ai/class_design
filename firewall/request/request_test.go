package request

import (
	"fmt"
	"testing"
)

func TestParseRequest(t *testing.T){
	type test struct{
		service_name string
		protocol string
		source_ip_address string
		destination_ip_address string
		source_port string
		destination_port string
	}
	tests := []struct{
		name string
		test test
	}{
		{
			name: "test1",
			test: test{"dns", "tcp", "XXX.XXX.XXX.XXX", "XXX.XXX.XXX.XXX", "80", "53"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
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
			req, err := ParseRequest(requestJSON)
			fmt.Println(req)
			if err != nil {
				t.Fatal(err)
			}
			if req.Name != want.service_name {
				t.Fatalf("Assersion Error: The value of service_name should be %s.", want.service_name)
			}
			if req.Protocol != want.protocol {
				t.Fatalf("Assersion Error: The value of protocol should be %s.", want.protocol)
			}
			if req.SourceIP != want.source_ip_address {
				t.Fatalf("Assersion Error: The value of source_ip_address should be %s.", want.source_ip_address)
			}
			if req.DestinationIP != want.destination_ip_address {
				t.Fatalf("Assersion Error: The value of destination_ip_address should be %s.", want.destination_ip_address)
			}
			if req.SourcePort != want.source_port {
				t.Fatalf("Assersion Error: The value of source_port should be %s.", want.source_port)
			}
			if req.DestinationPort != want.destination_port {
				t.Fatalf("Assersion Error: The value of destination_port should be %s.", want.destination_port)
			}
		})
	}
}
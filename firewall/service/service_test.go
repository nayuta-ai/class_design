package service

import (
	"testing"
)
func TestCreateService(t *testing.T){
	type port struct {
		protocol string
		port string
	}
	type test struct{
		service string
		port []port
	}
	var p []port
	p = append(p, port{"tcp", "53"})
	p = append(p, port{"udp", "53"})
	tests := []struct{
		name string
		test string
		want test
	}{
		{
			name: "test1",
			test: "dns",
			want: test{
				"DNS", p,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			service, err := CreateService(tt.test)
			if err != nil {
				t.Fatal(err)
			}
			if service.Short != tt.want.service {
				t.Fatalf("Assersion Error: The name of service should be %s.", tt.want.service)
			}
			for i, port := range service.Port {
				if string(port.Port) != tt.want.port[i].port {
					t.Fatalf("Assersion Error: The port of the service should be %s.", tt.want.port[i].port)
				}
				if string(port.Protocol) != tt.want.port[i].protocol{
					t.Fatalf("Assersion Error: The protocol of the service should be %s.", tt.want.port[i].protocol)
				}
			}
		})
	}
}
package request

import (
	"encoding/json"
)

type Request struct {
	Name            string `json:"service_name"`
	Protocol        string `json:"protocol"`
	SourceIP        string `json:"source_ip_address"`
	DestinationIP   string `json:"destination_ip_address"`
	SourcePort      string `json:"source_port"`
	DestinationPort string `json:"destination_port"`
}

func ParseRequest(requestJSON string) (Request, error) {
	var request Request
	err := json.Unmarshal([]byte(requestJSON), &request)
	if err != nil {
		return Request{}, err
	}
	return request, nil
}

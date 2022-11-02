package app

import (
	"log"
	"firewall/request"
	"firewall/service"
)

var NameService = make(map[string]service.Service)

func filter(request request.Request) bool {
	serviceName := request.Name
	_, ok := NameService[serviceName]
	if !ok {
		log.Printf("%s doesn't exist", serviceName)
		return false
	}
	for _, port := range NameService[serviceName].Port {
		if string(port.Port) == request.DestinationPort && string(port.Protocol) == request.Protocol {
			return true
		}
	}
	return false
}

func AddService(name string) error {
	service, err := service.CreateService(name)
	if err != nil {
		return err
	}
	NameService[name] = service
	return nil
}

func deleteService(name string) {
	delete(NameService, name)
}

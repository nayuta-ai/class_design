package service

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Service struct {
	XMLName xml.Name `xml:"service"`
	Short   string   `xml:"short"`
	Port    []Port   `xml:"port"`
}

type Port struct {
	XMLName  xml.Name `xml:"port"`
	Port     string   `xml:"port,attr"`
	Protocol string   `xml:"protocol,attr"`
}

var ServicesPath = "/go/src/firewall/services"

func CreateService(name string) (Service, error) {
	xmlFile, err := os.Open(ServicesPath + "/" + name)
	if err != nil {
		return Service{}, err
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var service Service
	err = xml.Unmarshal(byteValue, &service)
	if err != nil {
		return Service{}, err
	}
	return service, nil
}

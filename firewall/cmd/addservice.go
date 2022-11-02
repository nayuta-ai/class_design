package cmd

import (
	"encoding/xml"
	"firewall/service"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmdAddService = &cobra.Command{
	Use:   "add-service",
	Short: "Add a service",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("service-name")
		ports, _ := cmd.Flags().GetStringArray("add-port")
		serviceName := strings.ToLower(name)
		err := AddService(serviceName, name, ports)
		if err != nil {
			log.Println("Error encoding XML to file: ", err)
		}
	},
}

func AddService(name string, short string, ports []string) error {
	portSet := []service.Port{}
	for _, p := range ports {
		s := strings.Split(string(p), "/")
		portSet = append(
			portSet, service.Port{
				Port:     s[0],
				Protocol: s[1]},
		)
	}
	ser := service.Service{
		Short: short,
		Port:  portSet,
	}
	xmlFile, err := os.Create(service.ServicesPath + "/" + fmt.Sprintf("%v.xml", name))
	if err != nil {
		return err
	}
	xmlFile.WriteString(xml.Header)
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&ser)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	cmdAddService.Flags().StringP("service-name", "n", "", "Add a service name")
	cmdAddService.Flags().StringArrayP("add-port", "p", nil, "Add port sets")
	cmdAddService.Flags().StringP("add-discription", "d", "", "Add a detail of discription")
	rootCmd.AddCommand(cmdAddService)
}

package cmd

import (
	"log"
	"fmt"
	"io/ioutil"
	"firewall/service"
	"github.com/spf13/cobra"
)


var cmdList = &cobra.Command{
	Use: "list",
	Short: "Return all services",
	Long: "",
	Run: func(cmd *cobra.Command, args []string){
		err := fetchList(cmd)
		if err != nil {
			log.Println(err)
		}
	},
}

func fetchList(cmd *cobra.Command) error {
	files, err := ioutil.ReadDir(service.ServicesPath)
	if err != nil {
		return err
	}
	for _, f := range files{
		service, err := service.CreateService(f.Name())
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Service Name: %v\n", service.Short)
		for _, port := range service.Port {
			fmt.Fprintf(cmd.OutOrStdout(), "  Port/Protocol: %v/%v\n", port.Port, port.Protocol)
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(cmdList)
}
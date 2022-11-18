package main

import (
	"fmt"
	"log"
	"os"
	"taxi/system"
)

func main() {
	logs := "13:50:08.245 0.0<LF>13:50:11.123 4.0<LF>13:50:12.125 10.2<LF>13:50:13.100 8.7<LF>"
	totalFare, err := system.CalcTotalFare(logs)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(totalFare)
	os.Exit(0)
}

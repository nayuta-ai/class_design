package main

import (
	"fmt"
	"taxi/system"
)

func main() {
	logs := "13:50:08.245 0.0<LF>13:50:11.123 4.0<LF>13:50:12.125 10.2<LF>13:50:13.100 8.7<LF>"
	totalFare, _ := system.CalcTotalFare(logs)
	fmt.Println(totalFare)
}

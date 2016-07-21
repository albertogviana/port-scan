package main

import (
	"./nmap"
	"fmt"
)

func main() {
	host := ""
	result := nmap.Run(host)

	fmt.Print(result)
	//
	//openPorts := grep(output)
	//
	//v := convertStringToInt(openPorts)
	//
	//openPortsConfiguration := []int{
	//	22,
	//	443,
	//	9999,
	//}
	//
	//expectedUnfound, unexpectedFound := analyseResults(openPortsConfiguration, v)
	//
	//message := message(expectedUnfound, unexpectedFound)
	//fmt.Print(message)

}

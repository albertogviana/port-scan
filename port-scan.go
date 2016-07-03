package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func run() string {
	command := exec.Command("/usr/local/bin/nmap", "aisstaging.vesseltracker.com")
	var stdout bytes.Buffer
	command.Stdout = &stdout
	var stderr bytes.Buffer
	command.Stderr = &stderr
	error := command.Run()
	stdOut := stdout.String()
	stdErr := stderr.String()

	if error != nil || stdErr != "" {
		if error != nil {
			log.Printf("Error with command.Run(): %v\n", error)
		}
		if stdErr != "" {
			log.Printf("Stderr:\n%s", stdErr)
		}
		log.Fatal("Exiting!")
	}

	return stdOut
}

func grep(output string) string {
	portsRegex := regexp.MustCompilePOSIX("^[0-9]*")
	return strings.Join(removeEmptyStrings(portsRegex.FindAllString(output, -1)), " ")
}

func removeEmptyStrings(strings []string) []string {
	var ports []string
	for _, value := range strings {
		if value != "" {
			ports = append(ports, value)
		}
	}

	return ports
}

func convertStringToInt(ports string) []int {
	portsSlice := strings.Fields(ports)
	portsInteger := []int{}
	for _, port := range portsSlice {
		portInt, err := strconv.Atoi(strings.TrimSpace(port))
		if err != nil {
			log.Println("It was not possible to convert the string %s to integer\n Error: %v", port, err)
		}
		portsInteger = append(portsInteger, portInt)
	}

	return portsInteger
}

func main() {
	output := run()
	openPorts := grep(output)
	v := convertStringToInt(openPorts)

	fmt.Println(v)
}

package nmap

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func Run(host string) string {
	command := exec.Command("nmap", host)
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

func Grep(output string) string {
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

func analyseResults(expectedPorts []int, foundPorts []int) ([]int, []int) {
	expectedUnfound := compare(expectedPorts, foundPorts)
	unexpectedFound := compare(foundPorts, expectedPorts)
	return expectedUnfound, unexpectedFound
}

func compare(X, Y []int) []int {
	counts := make(map[int]int)
	var total int
	for _, value := range X {
		counts[value]++
		total++
	}

	for _, value := range Y {
		if _, ok := counts[value]; ok {
			counts[value]--
			total--
			if counts[value] <= 0 {
				delete(counts, value)
			}
		}
	}

	difference := make([]int, total)
	i := 0
	for value, count := range counts {
		for j := 0; j < count; j++ {
			difference[i] = value
			i++
		}
	}
	return difference
}

func message(expectedUnfoundPorts []int, unexpectedFoundPorts []int) string {
	var message string

	if len(expectedUnfoundPorts) > 0 {
		message += fmt.Sprintf(
			"The following ports were found filtered but were expected to be unfiltered:\n%d.\n\n",
			expectedUnfoundPorts,
		)
	}

	if len(unexpectedFoundPorts) > 0 {
		message += fmt.Sprintf(
			"The following ports were found unfiltered and are not part of the expected set:\n%d.\n\n",
			unexpectedFoundPorts,
		)
	}

	return message
}

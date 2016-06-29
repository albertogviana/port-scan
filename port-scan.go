package main

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func run() string {
	command := exec.Command("/usr/local/bin/nmap", "my-hosts")
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
	foundPortsString := strings.Join(removeEmptyStrings(portsRegex.FindAllString(output, -1)), " ")

	return foundPortsString
}

func removeEmptyStrings(strings []string) []string {
	var newStrings []string
	for _, value := range strings {
		if value != "" {
			newStrings = append(newStrings, value)
		}
	}

	return newStrings
}

func main() {
	output := run()
	grep(output)
}

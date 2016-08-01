package app

import (
	"fmt"
	"github.com/albertogviana/port_scan/config"
	"github.com/albertogviana/port_scan/message"
	"github.com/albertogviana/port_scan/nmap"
	"github.com/nlopes/slack"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Run the application
func Run(config *config.Config) {
	handler(config)
}

func handler(config *config.Config) {
	wg := new(sync.WaitGroup)
	for {

		for _, host := range config.Hosts {
			wg.Add(1)
			go checkPorts(host, config.Slack, wg)
		}
		wg.Wait()
		time.Sleep(time.Second * time.Duration(config.WaitNextCheck))
	}
}

func checkPorts(host config.HostConfiguration, config config.SlackConfiguration, wg *sync.WaitGroup) {
	defer wg.Done()
	ports := nmap.Run(host.Hostname)
	foundPorts := nmap.ConvertStringToInt(nmap.Parse(ports))

	expectedPorts := convertStringToInt(host.Port)

	expectedUnfound, unexpectedFound := nmap.AnalyseResults(expectedPorts, foundPorts)
	message := message.BuildMessage(host.Hostname, expectedUnfound, unexpectedFound)
	fmt.Print(message)

	if len(message) > 0 {
		sendMessage(message, config)
	}
}

func convertStringToInt(ports []string) []int {
	portsInteger := []int{}
	for _, port := range ports {
		portInt, err := strconv.Atoi(strings.TrimSpace(port))
		if err != nil {
			log.Println("It was not possible to convert the string %s to integer\n Error: %v", port, err)
		}
		portsInteger = append(portsInteger, portInt)
	}

	return portsInteger
}

func sendMessage(message string, config config.SlackConfiguration) {
	api := slack.New(config.Token)
	params := slack.PostMessageParameters{}
	params.Username = config.Username
	params.IconEmoji = config.IconEmoji

	channelID, timestamp, err := api.PostMessage(config.Channel, message, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

package nmap

import "fmt"

var expectedUnfoundPortsMsg = "The following ports were found filtered but were expected to be unfiltered: %d.\n"
var unexpectedFoundPortsMsg = "The following ports were found unfiltered and are not part of the expected set: %d.\n"

// BuildMessage builds the message with the results
func BuildMessage(expectedUnfoundPorts []int, unexpectedFoundPorts []int) string {
	var message string

	if len(expectedUnfoundPorts) > 0 {
		message += fmt.Sprintf(
			expectedUnfoundPortsMsg,
			expectedUnfoundPorts,
		)
	}

	if len(unexpectedFoundPorts) > 0 {
		message += fmt.Sprintf(
			unexpectedFoundPortsMsg,
			unexpectedFoundPorts,
		)
	}

	return message
}

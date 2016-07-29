package message

import "fmt"

var expectedUnfoundPortsMsg = "In the server %s the following ports were found filtered but were expected to be unfiltered: %d.\n"
var unexpectedFoundPortsMsg = "In the server %s the following ports were found unfiltered and are not part of the expected set: %d.\n"

// BuildMessage builds the message with the results
func BuildMessage(hostname string, expectedUnfoundPorts []int, unexpectedFoundPorts []int) string {
	var message string

	if len(expectedUnfoundPorts) > 0 {
		message += fmt.Sprintf(
			expectedUnfoundPortsMsg,
			hostname,
			expectedUnfoundPorts,
		)
	}

	if len(unexpectedFoundPorts) > 0 {
		message += fmt.Sprintf(
			unexpectedFoundPortsMsg,
			hostname,
			unexpectedFoundPorts,
		)
	}

	return message
}

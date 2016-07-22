package nmap

import (
	"fmt"
	"testing"
)

type testMessage struct {
	expectedUnfoundPorts []int
	unexpectedFoundPorts []int
	output               string
}

var buildMessageTests = []testMessage{
	{[]int{22, 80, 443}, []int{1026, 5000}, "The following ports were found filtered but were expected to be unfiltered: [22 80 443].\nThe following ports were found unfiltered and are not part of the expected set: [1026 5000].\n"},
	{[]int{22, 80, 443}, []int{}, "The following ports were found filtered but were expected to be unfiltered: [22 80 443].\n"},
	{[]int{}, []int{1026, 5000}, "The following ports were found unfiltered and are not part of the expected set: [1026 5000].\n"},
	{[]int{}, []int{}, ""},
}

func TestBuildMessage(t *testing.T) {
	for _, test := range buildMessageTests {
		message := BuildMessage(test.expectedUnfoundPorts, test.unexpectedFoundPorts)
		if fmt.Sprintf(test.output) != message {
			t.Error(
				"For: \n", test.output,
				"got: \n", message,
			)
		}
	}
}

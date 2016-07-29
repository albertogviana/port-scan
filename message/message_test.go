package message

import (
	"fmt"
	"testing"
)

type testMessage struct {
	hostname             string
	expectedUnfoundPorts []int
	unexpectedFoundPorts []int
	output               string
}

var buildMessageTests = []testMessage{
	{"host01.example.com", []int{22, 80, 443}, []int{1026, 5000}, "In the server host01.example.com the following ports were found filtered but were expected to be unfiltered: [22 80 443].\nIn the server host01.example.com the following ports were found unfiltered and are not part of the expected set: [1026 5000].\n"},
	{"host01.example.com", []int{22, 80, 443}, []int{}, "In the server host01.example.com the following ports were found filtered but were expected to be unfiltered: [22 80 443].\n"},
	{"host01.example.com", []int{}, []int{1026, 5000}, "In the server host01.example.com the following ports were found unfiltered and are not part of the expected set: [1026 5000].\n"},
	{"", []int{}, []int{}, ""},
}

func TestBuildMessage(t *testing.T) {
	for _, test := range buildMessageTests {
		message := BuildMessage(test.hostname,test.expectedUnfoundPorts, test.unexpectedFoundPorts)
		if fmt.Sprintf(test.output) != message {
			t.Error(
				"For: \n", test.output,
				"got: \n", message,
			)
		}
	}
}

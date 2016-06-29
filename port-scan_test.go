package main

import (
	"fmt"
	"testing"
)

type testGrep struct {
	input  string
	output string
}

var grepTests = []testGrep{
	{"22/tcp open  ssh", "22"},
	{"80/tcp open  http", "80"},
	{"\n100hdsajkds\nlkhjfd hhegdsk\n9\rsld", "100 9"},
}

func TestGrep(t *testing.T) {
	for _, test := range grepTests {
		output := grep(test.input)
		if test.output != output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", output,
			)
		}
	}
}

type testRemoveEmptyStrings struct {
	input  []string
	output []string
}

var removeEmptyStringsTests = []testRemoveEmptyStrings{
	{[]string{"", "22", "80", ""}, []string{"22", "80"}},
	{[]string{"", "22", "", "80", "", ""}, []string{"22", "80"}},
}

func TestRemoveEmptyStrings(t *testing.T) {
	for _, test := range removeEmptyStringsTests {
		output := removeEmptyStrings(test.input)
		if fmt.Sprintf("%v", test.output) != fmt.Sprintf("%v", output) {
			t.Error(
				"For", fmt.Sprintf("%v", test.input),
				"expected", fmt.Sprintf("%v", test.output),
				"got", fmt.Sprintf("%v", output),
			)
		}
	}
}

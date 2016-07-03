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

type testStringToIntSlice struct {
	str  string
	ints []int
}

var stringToIntTests = []testStringToIntSlice{
	{"22 23 25 80 111 443 465 587", []int{22, 23, 25, 80, 111, 443, 465, 587}},
	{"789", []int{789}},
	{"789 ", []int{789}},
	{" 789 ", []int{789}},
	{" 789", []int{789}},
	{"5", []int{5}},
}

func TestConvertStringToInt(t *testing.T) {
	for _, test := range stringToIntTests {
		converted := convertStringToInt(test.str)
		convertedString := fmt.Sprintf("%d", converted)
		testString := fmt.Sprintf("%d", test.ints)
		if convertedString != testString {
			t.Error(
				"For", test.str,
				"expected", test.ints,
				"got", converted,
			)
		}
	}
}

package nmap

import (
	"fmt"
	"testing"
)

type testParse struct {
	input  string
	output string
}

var parseTests = []testParse{
	{"22/tcp open  ssh", "22"},
	{"80/tcp open  http", "80"},
	{"\n100hdsajkds\nlkhjfd hhegdsk\n9\rsld", "100 9"},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		output := Parse(test.input)
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
		converted := ConvertStringToInt(test.str)
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

// compare
type testComparePorts struct {
	expected        []int
	found           []int
	expectedUnfound []int
	unexpectedFound []int
}

//// expected, found, expectedUnfound, unexpectedFound
var comparePortsTests = []testComparePorts{
	{[]int{80, 443, 5432}, []int{80, 443}, []int{5432}, []int{}},
	{[]int{80, 443, 5432}, []int{80, 443, 9876}, []int{5432}, []int{9876}},

	{[]int{80, 443}, []int{80, 443, 5432}, []int{}, []int{5432}},
	{[]int{443, 80}, []int{80, 443, 5432}, []int{}, []int{5432}},
	{[]int{443, 80}, []int{5432, 443, 80}, []int{}, []int{5432}},

	{[]int{80, 443}, []int{443, 80}, []int{}, []int{}},
	{[]int{80, 443}, []int{80, 443}, []int{}, []int{}},
}

func TestComparePorts(t *testing.T) {
	for _, test := range comparePortsTests {
		expectedUnfound, unexpectedFound := AnalyseResults(test.expected, test.found)
		foundEuf := fmt.Sprintf("%v", expectedUnfound) != fmt.Sprintf("%v", test.expectedUnfound)
		foundUef := fmt.Sprintf("%v", unexpectedFound) != fmt.Sprintf("%v", test.unexpectedFound)
		if foundEuf || foundUef {
			t.Error(
				"For", test.expected, "and", test.found,
				"expected", test.expectedUnfound, "and", test.unexpectedFound,
				"got", expectedUnfound, "and", unexpectedFound,
			)
		}
	}
}

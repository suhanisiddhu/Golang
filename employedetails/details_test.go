package employedetails

import (
	"testing"
)

type testCases struct {
	desc  string
	input employee
	out   output
}

func TestDetails(t *testing.T) {
	s := []testCases{
		{">=22", employee{2, "chaithra", 27}, output{true, employee{2, "chaithra", 27}}},
		//	{"<22", employee{}, output{true, employee{}}},
	}

	for _, v := range s {
		res := details(v.input)
		if res != v.out {
			t.Errorf("failed")

		}

	}
}

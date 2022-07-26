package Employee

import (
	"testing"
)

func TestEmployee(t *testing.T) {
	s := []Employee{
		{"neha", "21 - 04 - 2000", 1, true},
		{"rita", "22 - 05 - 1999", 9, true},
		{"   ", "", 0, false},
		{"", "12-06-2002", -3, false},
		{"gita", "  ", 0, false},
		{"roli", "30-04-2001", 14, true},
		{"ram", "20-05-1998", -1, false},
		{"abcd", "", 14, false},
	}
	for _, v := range s {
		res := v.printdetail()
		if res == v.out {
		} else {
			t.Errorf("failed")

		}

	}

}

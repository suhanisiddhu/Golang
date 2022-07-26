package primeTest

import (
	"testing"
)

type TestCases struct {
	desc   string
	input  int
	output bool
}

func TestPrime(t *testing.T) {
	testCases := []TestCases{

		{"prime", 17, true},
		{"noPrime", -20, false},
		{desc: "noPrime", input: 0, output: false},
		{desc: "noPrime", input: 1, output: false},
	}
	for _, tc := range testCases {

		res := primeNum(tc.input)
		if res == tc.output {

		} else {
			t.Errorf("case failed ")
		}
	}
}
func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNum(i)
	}
}

/*func TestPrime(t *testing.T) {
	input := 42
	output := false
	res := primeNum(input)
	if res == output {
		fmt.Println("case passed ")
	} else {
		t.Errorf("case failed ")
	}
}*/

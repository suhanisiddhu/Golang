package twofer

import (
	"fmt"
	"testing"
)

func TestTwofer(t *testing.T) {
	s := []struct {
		desc   string
		input  string
		output bool
	}{
		{"twofer", "bob", true},
		{"twofer", " ", false},
	}

	for _, v := range s {
		res := name(v.input)
		if res == v.output {
			fmt.Println("case passed")
		} else {
			t.Errorf("failed")

		}

	}
}
func BenchMarkTwoFer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name("suhani")
	}
}

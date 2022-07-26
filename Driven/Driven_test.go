package Driven

import (
	"fmt"
	"testing"
)

func TestDriven(t *testing.T) {
	s := []struct {
		input  string
		output string
	}{
		{"How are you?", "Sure"},
		{"YELL AT HIM", "Whoa, chill out!"},
		{"HOW ARE YOU?", "Calm down, I know what I'm doing!"},
		{"        ", "Fine. Be that way!"},
		{"anything else", "Whatever"},
	}
	for _, v := range s {
		res := drive(v.input)
		if res == v.output {
			fmt.Println("case passed")
		} else {
			t.Errorf("failed")

		}

	}
}

func BenchMarkTwoFer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drive("SUHANI")
	}
}

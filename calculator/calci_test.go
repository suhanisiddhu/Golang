package calculator

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	s := []struct {
		description string
		x           int
		y           int
		operator    string
		output      int
	}{
		{"addition", 4, 6, "+", 10},
		{"subtraction", 8, 4, "-", 4},
		{"multiplication", 6, 2, "*", 12},
		{"divide", 12, 2, "/", 6},
		{"notavalidoperator", 20, 5, "&", -1},
		{"divide", 20, 0, "/", -1},
		{"addition", 30, 12, "", -1},
	}

	for _, v := range s {
		res := calculate(v.x, v.y, v.operator)
		if res == v.output {
		} else {
			t.Errorf("failed")

		}

	}
}

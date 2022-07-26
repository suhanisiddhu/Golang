package Interface

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape  Shape
		output float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Square{4}, 16.0},
		{Circle{10}, math.Pi * 100},
		{Triangle{12, 3, 5, 6}, 15.0},
	}

	for _, tt := range areaTests {
		v := tt.shape.Area()
		if v != tt.output {
			t.Errorf("case failed")
		}
	}

	perimeterTests := []struct {
		shape  Shape
		output float64
	}{
		{Rectangle{12, 6}, 36.0},
		{Square{4}, 16.0},
		{Circle{10}, 20 * math.Pi},
		{Triangle{12, 3, 5, 6}, 20.0},
	}

	for _, x := range perimeterTests {
		y := x.shape.Perimeter()
		if y != x.output {
			t.Errorf("case failed")
		}

	}
}

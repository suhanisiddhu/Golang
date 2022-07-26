package primeTest

import (
	"math"
)

func primeNum(num int) bool {
	if num < 2 {

		return false
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {

			return false
		}
	}

	return true
}

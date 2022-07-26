package main

import "fmt"

func main() {

	s1 := sum(1, 2, 3)
	s2 := sum(1, 2, 3, 4)
	s3 := sum(1, 2, 13, 4, 5, 6)

	fmt.Println(s1, s2, s3)
}

func sum(nums ...int) int {

	res := 0

	for _, n := range nums {
		res += n
	}

	return res
}

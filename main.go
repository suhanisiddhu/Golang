package main

import (
	"Time/time/calculatetime"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(calculatetime.Time(t))
	//a := 10
	//b := 20
	//fmt.Println(calculatArea.Area(a, b))
}

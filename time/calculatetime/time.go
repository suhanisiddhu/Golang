package calculatetime

import (
	"math"
	"time"
)

func Time(s time.Time) time.Time {
	return s.Add(time.Duration(math.Pow10(9)) * time.Second)
}

package bflowcontrol

import (
	"fmt"
	"math"
)

func SqrtIf(x float64) string {
	if x < 0 {
		return SqrtIf(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

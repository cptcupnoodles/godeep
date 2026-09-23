package abasic

import "fmt"

func Convertion() {

	var x int = 10
	var y float64 = 8.5
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)
}

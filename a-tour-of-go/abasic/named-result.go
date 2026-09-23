package abasic

import "fmt"

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func NamedResult() {
	fmt.Println(Split(17))
}

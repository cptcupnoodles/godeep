package abasic

import "fmt"

func Swap(x, y string) (string, string) {
	return y, x
}

func MultiResults() {
	a, b := Swap("PING", "PONG")
	fmt.Println(a, b)
}

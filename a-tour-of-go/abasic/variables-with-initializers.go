package abasic

import "fmt"

var i, j int = 1, 2

func VarsWithInitializers() {
	var l, k, h = true, false, "!yes"
	fmt.Println(i, j, l, k, h)
}

package abasic

import "fmt"

func ShortVariableDeclarations() {
	var i, j int = 1, 2
	k := 3
	java, golang, python := "very-good", "top-tier", "meh"
	fmt.Println(i, j, k, java, golang, python)
}

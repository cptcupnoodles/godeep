package cmoretypes

import "fmt"

func Pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	fmt.Println(j)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

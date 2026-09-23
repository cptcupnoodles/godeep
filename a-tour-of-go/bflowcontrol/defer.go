package bflowcontrol

import "fmt"

func DeferKey() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

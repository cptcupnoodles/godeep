package bflowcontrol

import "fmt"

func ForLoop() {
	sum := 0
	for i := range 10 {
		sum += i
	}
	fmt.Println(sum)
}

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
//

func fibonacci() func() int {
	current := 0
	last := 0
	return func() int {
		if current == 0 {
			current++
			return last
		}
		current += last
		last = current - last

		return last
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

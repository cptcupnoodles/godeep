package main

import "fmt"

func countdown(start int) (sum int, valid bool) {
	fmt.Println("Starting countdown")

	defer fmt.Println("Countdown finished")

	if start < 1 {
		fmt.Println("Invalid start")
		return 0, false
	}

	for i := start; i >= 1; i-- {
		fmt.Println("Now", i)
		sum += i
		defer fmt.Println("Saved:", i)
	}
	fmt.Println("Sum:", sum)

	return sum, true
}

func main() {
	fmt.Println(countdown(6))
}

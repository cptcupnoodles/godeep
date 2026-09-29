package main

import "fmt"

func apply(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))

	for i, number := range numbers {
		result[i] = fn(number)
	}

	return result
}

func double(number int) int {
	return number * 2
}

func square(number int) int {
	return number * number
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8}

	double := apply(n, double)
	square := apply(n, square)

	fmt.Println("Doubled: ", double)
	fmt.Println("Squared: ", square)
}

package main

import (
	"fmt"
)

type Bin struct {
	Label string
	Units int
}

var bins = [5]Bin{
	{Label: "A", Units: 10},
	{Label: "B", Units: 20},
	{Label: "C", Units: 30},
	{Label: "D", Units: 40},
	{Label: "E", Units: 50},
}

var left = bins[0:3]
var right = bins[2:]
var empty = right[:0]

func transfer(bins []Bin, from int, to int, amount int) bool {

	if len(bins) == 0 {
		fmt.Println("Empty slice")
		return false
	}
	if from < 0 || to < 0 {
		fmt.Println("Index is negative")
		return false
	}

	if from > len(bins)-1 || to > len(bins)-1 {
		fmt.Println("Index out of range")
		return false
	}

	if from == to {
		fmt.Println("It cannot be the same index")
		return false
	}

	if amount <= 0 {
		fmt.Println("Amount is zero")
		return false
	}

	if amount > bins[from].Units {
		fmt.Println("Source has insufficient units.")
		return false
	}

	source, destination := &bins[from].Units, &bins[to].Units

	*source -= amount
	*destination += amount

	return true
}

func totalUnits(bins []Bin) int {
	total := 0
	for i := 0; i < len(bins); i++ {
		total += bins[i].Units
	}

	return total
}

func main() {
	fmt.Println(bins)
	fmt.Println(transfer(left, 0, 2, 4))
	fmt.Println(bins)
	fmt.Println(transfer(right, 0, 2, 10))
	fmt.Println(totalUnits(bins[:]))
	fmt.Println(bins)
	fmt.Println(len(left), cap(left), len(right), cap(right))
	fmt.Println(len(empty), cap(empty))

}

// 1.second transfer change the value becuase right[0]
// share the same address as left[2] that is why changes
// on those index will affect both slices.
//
// 2.index check must use the length of the slice because length derived from the real array not the capacity of the array, a slice can have an n capacity but can have zero length so using length will be accurate to be use for checking.
//
// 3.empty has zero lenght becuase it derived from the start of the slice before index 0 which is nothing.

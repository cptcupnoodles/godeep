package cmoretypes

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Range() {
	rows := make([][]int, 2)

	fmt.Println(rows)
	// i is the index and v is the value of the index
	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	// only the value
	// for _, v := range pow {
	// 	fmt.Println(v)
	// }

	// only the index
	// for i := range pow {
	//    fmt.Println(i)
	// }

	// v is a copy of the element. Changing v does not change the slice:

	// To change the slice, use the index:
	// for i := range pow {
	//    pow[i] = 0
	// }
}

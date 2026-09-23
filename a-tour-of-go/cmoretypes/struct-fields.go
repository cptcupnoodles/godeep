package cmoretypes

import "fmt"

func StructFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

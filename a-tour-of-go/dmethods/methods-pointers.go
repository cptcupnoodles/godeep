package dmethods

import (
	"fmt"
)

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MethodsPointers() {
	v := Vertex{3, 4}
	v.Scale(1)
	fmt.Println(v.Abs())
}

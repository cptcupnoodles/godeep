package dmethods

import "fmt"

func EmptyInterface() {
	var i interface{}
	describeThis(i)

	i = 42
	describeThis(i)

	i = "hello"
	describeThis(i)
}

func describeThis(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

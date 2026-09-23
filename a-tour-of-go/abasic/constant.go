package abasic

import "fmt"

const Pi = 30.6

func Constant() {
	const World = "Earth"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

package bflowcontrol

import (
	"fmt"
	"time"
)

func SwitchEvalOrder() {
	today := time.Now().Weekday()
	fmt.Println("Today is ", today)
	fmt.Println("When's Monday?")

	switch time.Monday {
	case today + 0:
		fmt.Println("Today. ")
	case (today + 1) % 7:
		fmt.Println("Tomorrow. ")
	case (today + 2) % 7:
		fmt.Println("In two days. ")
	default:
		fmt.Println("Too far away.")
	}
}

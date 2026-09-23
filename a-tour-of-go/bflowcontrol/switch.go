package bflowcontrol

import (
	"fmt"
	"runtime"
)

func SwitchCases() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS. ")
	case "linux":
		fmt.Println("Linux. ")
	default:
		fmt.Printf("%s. \n", os)
	}
}

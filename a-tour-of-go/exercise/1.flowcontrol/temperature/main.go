package main

import "fmt"

func converter(celsius float64) (float64, float64) {
	const f = float64(32)
	const k = 273.15
	fahrenheit := celsius*9/5 + f
	kelvin := celsius + k

	return fahrenheit, kelvin
}

func main() {
	var temp int = -10
	f, k := converter(float64(temp))

	var loc string = "Tacurong City"
	var belowZero = temp < 0
	fmt.Println("| Location | Celsius | Fahrentheit | Kelvin | Below-Zero |\n", "| ", loc, " | ", temp, " | ", f, " | ", k, " | ", belowZero, " |")
}

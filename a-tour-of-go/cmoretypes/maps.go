package cmoretypes

import "fmt"

type Geo struct {
	Lat, Long float64
}

var m map[string]Geo

func Maps() {
	m = make(map[string]Geo)
	m["Bell Labs"] = Geo{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

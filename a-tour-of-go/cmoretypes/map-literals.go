package cmoretypes

import "fmt"

var ml = map[string]Geo{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func MapLiteral() {
	fmt.Println(ml)
}

package main

import (
	"fmt"
)

func evaluateScore(score int) (grade string, passed bool) {
	if score < 0 || score > 100 {
		return "Invalid", false
	}

	const passingScore = 60

	if score >= passingScore {
		passed = true
	} else {
		passed = false
	}

	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	return

}

func message(grade string) string {
	switch grade {
	case "A":
		return "Excellent"
	case "B":
		return "Very Good"
	case "C":
		return "Good"
	case "D":
		return "Bad"
	case "F":
		return "Very Bad"
	default:
		return "Invalid"
	}
}

func main() {
	s := 99
	if g, p := evaluateScore(s); p {
		fmt.Println(s, p, g, message(g))
	} else {
		fmt.Println(g, p)
	}
}

// ### Explain your result
// Well testing >= 60 will result into highest grade which is incorrect because >= 60 will always return A.

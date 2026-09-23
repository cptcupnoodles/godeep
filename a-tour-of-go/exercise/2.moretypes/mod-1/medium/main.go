package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

var students = [5]Student{
	{Name: "Ana", Score: 55},
	{Name: "Ben", Score: 58},
	{Name: "Cara", Score: 76},
	{Name: "Dan", Score: 98},
	{Name: "Eli", Score: 40},
}

var group = students[1:4]

func addBonus(students []Student, bonus int) bool {
	if bonus < 0 {
		return false
	}

	for i := 0; i < len(students); i++ {
		var s = &students[i]
		s.Score += bonus

		if s.Score > 100 {
			s.Score = 100
		}
	}
	return true
}

func summarize(students []Student) (total int, passed int) {
	sum := 0
	t := 0
	for i := 0; i < len(students); i++ {
		if students[i].Score >= 60 {
			sum += students[i].Score
			t++
		}
	}

	return sum, t
}

func main() {
	fmt.Println(addBonus(group, 5))
	fmt.Println(len(group), cap(group))
	fmt.Println(summarize(students[:]))
}

// - addBonus can change array elements becuase data passed into it is pointed into the origin data
// - it will not change the data in the origin becuase what you change is just a copy not the pointer that is pointed into the origin data
//

package main

import (
	"fmt"
	"log"
	"strconv"
)

func calculate(array []int) float64 {
	var sum int = 0
	for _, score := range array {
		sum += score
	}
	average := float64(sum) / float64((len(array)))
	return average
}

func main() {
	students := [][2]string{
		{"Alice", "80"},
		{"Bob", "90"},
		{"Charlie", "85"},
		{"Diana", "95"},
	}

	var students_score []int
	for _, student := range students {
		score, err := strconv.Atoi(student[1])
		if err != nil {
			log.Fatal(err)
		}

		students_score = append(students_score, score)
	}

	a := calculate(students_score)
	fmt.Print(a)
}

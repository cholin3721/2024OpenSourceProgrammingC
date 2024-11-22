package main

import (
	"fmt"
)

func main() {
	var emptySlice []bool
	emptySlice = make([]bool, 5)
	fmt.Printf("#%#v\n", emptySlice)

	var gpa [5]float64 = [5]float64{3, 5, 4.1, 3.5, 4.23}
	gpa_slice := gpa[1:4]
	//gpa_slice[1] = 2.71
	gpa[2] = 2.71
	//gpa_slice = append(gpa_slice, 4.3)
	gpa_slice = append(gpa_slice, 4.3, 5.55)
	//이 순간 메모리 위치가 달라짐
	fmt.Println(len(gpa_slice), gpa_slice, gpa)
}

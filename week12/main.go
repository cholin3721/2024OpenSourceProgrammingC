package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int
	scores[1] = 90
	//fmt.Println(scores[1], scores[0]) //, scores[3])
	//fmt.Printf("%#v\n", scores)
	//fmt.Println(scores)

	for i := 0; i < len(scores); i++ { //unsafe
		fmt.Printf("%d ", scores[i])
	}

	//리터럴은 값 그자체를 의미함
	//var dates [3]time.Time
	//dates[1] = time.Unix(1947200001, 0)
	//fmt.Println(dates[1])

	//dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)}
	//fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1947200001, 0), //need comma, 대괄호를 붙이면 가능하긴함
	}
	//fmt.Printf("%#v\n", dates) \n<<rune 값 : 10
	//fmt.Println(dates)
	//fmt.Println(dates[0], dates[1], dates[2])
	//fmt.Println()
	//for i, date := range dates {
	//	fmt.Println(i, date)
	//}
	fmt.Println()
	for _, date := range dates {
		fmt.Println(date)
	}

}

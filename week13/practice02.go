package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func keyboard() string {
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(data)
}

func toInt(data string) (int, bool) {
	score, err := strconv.Atoi(data)
	if err != nil {
		return 0, false
	}

	return score, true
}

func calculate1(array []int) float64 {
	var sum int = 0
	for _, score := range array {
		sum += score
	}
	average := float64(sum) / float64((len(array)))
	return average
}

func main() {
	var students_name []string
	var students_score []int

	for {
		fmt.Printf("1. 학생 추가\n2. 학생 점수 조회\n3. 전체 학생 점수 출력\n4. 평균 점수 계산\n5. 종료\n")
		choice := keyboard()
		switch choice {
		case "1":
			fmt.Print("Enter the Name : ")
			name := keyboard()
			students_name = append(students_name, name)
			fmt.Print("Enter the Score : ")
			score, value := toInt(keyboard())
			if value == true {
				students_score = append(students_score, score)
			} else {
				score = 0
				students_score = append(students_score, score)
			}
			continue

		case "2":
			if len(students_name) == 0 {
				fmt.Println("입력된 학생이 없습니다.")
				continue
			}
			fmt.Print("Enter the Name : ")
			name := keyboard()
			match := false
			for index, n := range students_name {
				if name == n {
					fmt.Println(n, ":", students_score[index])
					match = true
					continue
				}
			}
			if match == false {
				fmt.Println("학생을 찾을 수 없습니다.")
				continue
			}

		case "3":
			if len(students_score) == 0 {
				fmt.Println("학생 점수가 없습니다.")
				continue
			}
			for index, name := range students_name {
				fmt.Println(name, ":", students_score[index])
			}
			continue

		case "4":
			if len(students_score) == 0 {
				fmt.Println("학생 점수가 없습니다.")
				continue
			}
			average := calculate1(students_score)
			fmt.Println("전체 학생 평균 : ", average)

		case "5":
			return

		default:
			fmt.Println("다시 입력해주세요 올바르지 않은 입력입니다.")
			continue

		}

	}

}

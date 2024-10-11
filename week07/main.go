package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("점수 입력 : ")
	i, err := in.ReadString('\n')
	//fmt.Println(i, err)
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)                //줄바꿈등 제거, 파이썬의 strip 함수와 비슷
	score, _ := strconv.ParseInt(i, 10, 32) //16진 정수형(32bit)으로 변환
	var grade string
	if score >= 90 {
		grade = "A"

	} else {
		grade = "BCDF"

	}
	fmt.Printf("%d점은 %s등급 입니다.\n", score, grade)
}

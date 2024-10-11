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
	fmt.Print("Input Your Score : ")
	i, err := in.ReadString('\n')
	//fmt.Println(i, err)
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)                //줄바꿈등 제거, 파이썬의 strip 함수와 비슷
	score, _ := strconv.ParseInt(i, 10, 32) //10진 정수형(32bit)으로 변환
	if score >= 90 {
		fmt.Println("A")
		fmt.Printf("%d", score)
	} else {
		fmt.Println("BCDF")
		fmt.Printf("%d", score)
	}

}

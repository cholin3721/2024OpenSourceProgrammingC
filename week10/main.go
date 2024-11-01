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
	fmt.Print("Input Number : ")
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	//counts := 0
	//bool타입은 1바이트, int는 4바이트라서 용량 더 적음
	//is 변수 는 플래그 변수 is a number?
	//bool타입으로 바꿔서 가독성을 높임 메모리 down
	var isPrime bool = true

	j := 2
	for j < n {
		if n%j == 0 {
			isPrime = false //더하기 연산자 제거
		}
		j++
	}

	if isPrime {
		fmt.Printf("%d is Prime Number", n)
	} else {
		fmt.Printf("%d is NOT Prime Number", n)
	}

}

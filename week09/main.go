package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //seed 값이 고정되면 난수가 고정됨, 유닉스 시간을 가져와서 시드값 고정 못하게 함
	answer := rand.Intn(6) + 1
	fmt.Println(answer)

	fmt.Print("숫자 입력 : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	//fmt.Println(i, err)
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)      //줄바꿈등 제거, 파이썬의 strip 함수와 비슷
	guess, err := strconv.Atoi(i) //16진 정수형(32bit)으로 변환
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)

	if answer < guess {
		fmt.Println("입력하신 수는 정답입니다.")
	} else if answer > guess {
		fmt.Println("입력하신 수가 정답보다 작은 수입니다. LOW")
	} else {
		fmt.Println("입력하신 수는 정답보다 큰 수입니다. HIGH")
	}
}

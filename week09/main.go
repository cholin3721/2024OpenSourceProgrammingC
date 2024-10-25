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
	var win bool = false

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력 : ", 3-guesses)
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

		if answer == guess {
			fmt.Println()
			fmt.Println("입력하신 수는 정답입니다. 선물은 없습니다.")
			win = true
			break
		} else if answer > guess {
			fmt.Println()
			fmt.Println("입력하신 수가 정답보다 작은 수입니다. LOW")
			fmt.Println()
		} else {
			fmt.Println()
			fmt.Println("입력하신 수는 정답보다 큰 수입니다. HIGH")
			fmt.Println()
		}

	}

	if win == true {
		fmt.Println("You Win!!")
	} else {
		fmt.Println("허접ㅋ")
	}
}

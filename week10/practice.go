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
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	judge := false
	for tryNum := 1; tryNum <= 3; tryNum++ {
		fmt.Printf("%d번째 시도 : ", tryNum)
		in := bufio.NewReader(os.Stdin)
		i, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}

		if guess == answer {
			judge = true
			break

		} else if guess > answer {
			fmt.Printf("숫자가 너무 커 임마 ㅋㅋ\n")
			fmt.Println()
		} else {
			fmt.Printf("숫자가 작다고 임마 ㅋㅋ\n")
			fmt.Println()
		}
	}
	if judge == false {
		fmt.Printf("병신ㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋ 정답은 %d야 ㅋㅋ\n", answer)
		fmt.Println()
	} else {
		fmt.Printf("올ㅋ")
		fmt.Println()
	}
}

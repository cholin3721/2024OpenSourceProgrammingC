package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Printf("%f\n", math.Sqrt(16.0))
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

	var isPrime bool = true

	if n <= 1 {
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 { //2를 제외한 짝수는 모두 소수가 아니다
		isPrime = false
	} else { //odd numbers
		j := 3 // start value
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", j) // check j loop
			j = j + 2            //increment

		}
	}
	if isPrime {
		fmt.Printf("%d is Prime Number", n)
	} else {
		fmt.Printf("%d is NOT Prime Number", n)
	}

}

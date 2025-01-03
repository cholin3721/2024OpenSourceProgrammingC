package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	//var isPrime bool = true

	if n <= 1 {
		return false
	} else if n == 2 {
		//isPrime = true
		return true
	} else if n%2 == 0 {
		//isPrime = false
		return false
	} else {
		j := 3
		//for j <= int(math.Sqrt(float64(n)))
		for j*j <= n {
			if n%j == 0 {
				//isPrime = false
				//break
				return false
			}
			fmt.Printf("%d ", j)
			j = j + 2

		}
	}
	return true

}
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

	if isPrime(n) {
		fmt.Printf("%d is Prime Number", n)
	} else {
		fmt.Printf("%d is NOT Prime Number", n)
	}

}

// 다른 언어로도 해볼것
// 집에서 3.1 버전으로 합성수 소수 판별 해볼 것

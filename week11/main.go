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
			//fmt.Printf("%d ", j)
			j = j + 2

		}
	}
	return true

}
func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input Start Number : ")
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Input End Number : ")
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n2, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j < n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}

}

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

	var isPrime bool = true
	//bug fix
	if n <= 1 { //A prime number is a champion number that is a number greater than 1 that is a divisor of 1 and itself.
		isPrime = false
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				isPrime = false
				break //performance up
			}
			//fmt.Printf("%d ", j) // check j loop
			j++

		}
	}
	if isPrime {
		fmt.Printf("%d is Prime Number", n)
	} else {
		fmt.Printf("%d is NOT Prime Number", n)
	}

}

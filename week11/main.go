package main

import (
	"fmt"
	"log"
	"week11/keyboard"
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

	fmt.Print("Input Start Number : ")
	n1, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Input End Number : ")
	n2, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j < n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(a int) bool {
	if a <= 1 {
		fmt.Print("Prime Number should bigger than 1")
		return false
	} else if a == 2 {
		return true
	} else if a%2 == 0 {
		return false
	} else {
		j := 3
		for j*j <= a {
			if a%j == 0 {
				return false
			}
			fmt.Printf("%d ", j)
			j = j + 2
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("Input the Number : ")
		in, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inst := strings.TrimSpace(in)
		num, err := strconv.Atoi(inst)
		if err != nil {
			log.Fatal(err)
		}

		judge := isPrime(num)

		if judge {
			fmt.Printf("%d is Prime Number\n", num)

		} else {
			fmt.Printf("%d is NOT Prime Number\n", num)
		}

		fmt.Println("Wanna Continue? (Y/N)")
		try, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		i := strings.TrimSpace(try)
		if i == "Y" || i == "y" {
			continue
		} else {
			break
		}
	}
}

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

	counts := 0

	j := 1
	for j <= n {
		if n%j == 0 {
			counts++
		}
		j++
	}

	if counts == 2 {
		fmt.Printf("%d is Prime Number", n)
	} else {
		fmt.Printf("%d is not Prime Number", n)
	}

}

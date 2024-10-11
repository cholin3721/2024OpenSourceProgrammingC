package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input Your name : ")
	name, err := in.ReadString('\n')
	//fmt.Println(i, err)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}

}

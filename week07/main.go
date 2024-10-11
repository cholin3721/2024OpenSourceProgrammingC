package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input Your name : ")
	name, err := in.ReadString('\n')
	//fmt.Println(i, err)
	fmt.Println(name)
	fmt.Println(err)
}

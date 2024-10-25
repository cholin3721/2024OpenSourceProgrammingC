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
	fmt.Print("너의 성적을 입력하거라 : ")
	rdyforinput := bufio.NewReader(os.Stdin)
	input, err := rdyforinput.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	edit := strings.TrimSpace(input)
	editedNum, err := strconv.ParseFloat(edit, 64)

	if err != nil {
		log.Fatal(err)
	}
	var status string
	if editedNum > 90 {
		status = "Good job!"
	} else if editedNum > 80 {
		status = "Okay"
	} else if editedNum > 70 {
		status = "Hmm, soso"
	} else {
		status = "병신 ㅋㅋ"
	}
	fmt.Println(status)
}

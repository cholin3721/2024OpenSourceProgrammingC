package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //seed 값이 고정되면 난수가 고정됨
	target := rand.Intn(6)
	fmt.Println(target)
}

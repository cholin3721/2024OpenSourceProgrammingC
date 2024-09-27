package main

import (
	"fmt" //#include <stdio.h>
	"math"
	"reflect"
	"strings"
)

func main() {
	//var i int = 55
	//var f float32 = 3.99
	//var i int
	//i = 55

	//i := 55
	i := "55"
	f := 3.99
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
	fmt.Println(f, math.Ceil(3.49))
	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("i는 %s\n", i)
	fmt.Print("i는 ", i, "\n")
	fmt.Println("i는", i)

}

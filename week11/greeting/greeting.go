package greeting //패키지는 소문자로 하고 폴더 이름이랑 똑같아야함

import "fmt"

func hi(name string) {
	fmt.Printf("Hi, %s!\n", name)
}

func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func EnglishGreetings(name string) {
	hello(name)
	hi(name)
}

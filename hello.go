package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello,"

func main() {
	fmt.Println(Hello("Golang"))
}

func Hello(name string) string {
	return englishHelloPrefix + name
}

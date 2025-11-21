package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"
const spanish = "Spanish"
const french = "French"

func main() {
	fmt.Println(Hello("Golang", "English"))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "Golang"
	}

	return greetingPrefix(language) + name

	/*if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		language = frenchHelloPrefix + name
	}
	return englishHelloPrefix + name*/
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

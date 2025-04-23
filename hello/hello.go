package main

import "fmt"

const (
	spanish                 = "Spanish"
	french                  = "French"
	englishHelloWorldPrefix = "Hello, "
	spanishHelloWorldPrefix = "Hola, "
	frenchHelloWorldPrefix  = "Bonjour, "
)

func Hello(name string, langauge string) string {

	if name == "" {
		name = "World"
	}

	languagePrefix := greetingPrefix(langauge)
	return languagePrefix + name
}

func greetingPrefix(langauge string) (prefix string) {
	switch langauge {
	case spanish:
		prefix = spanishHelloWorldPrefix
	case french:
		prefix = frenchHelloWorldPrefix
	default:
		prefix = englishHelloWorldPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Amit", ""))
}

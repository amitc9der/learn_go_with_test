package main

import "fmt"

const englishHelloWorldPrefix = "Hello, "

func Hello(name string) string {

	if name == "" {
		name = "World"
	}

	return englishHelloWorldPrefix + name
}

func main() {
	fmt.Println(Hello("Amit"))
}

package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func runInterfaces() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// receiver variable property not needed in method signature as not used inside
// type only here to explain interfaces in this example
func (englishBot) getGreeting() string {

	return "Hello There"
}

func (spanishBot) getGreeting() string {

	return "Hola!"
}

func printGreeting(b bot) {

	fmt.Println(b.getGreeting())
}

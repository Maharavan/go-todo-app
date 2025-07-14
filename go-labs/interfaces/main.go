package main

import (
	"fmt"
)

type englishBot struct{ message string }
type spanishBot struct{ message string }
type brazilBot struct{ message string }
type bot interface {
	getGreetings() string
}

func main() {
	eb := englishBot{message: "Hello"}
	sb := spanishBot{message: "Hola"}
	bb := brazilBot{message: "bem vindo"}
	printGreeting(eb)
	printGreeting(sb)
	printGreeting(bb)
}

func (eb englishBot) getGreetings() string {
	return eb.message
}

func (bb brazilBot) getGreetings() string {
	return bb.message
}
func (sb spanishBot) getGreetings() string {
	return sb.message
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

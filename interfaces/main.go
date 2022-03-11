package main

import "fmt"

type bot interface {
	GetGreeting() string
	GetGoodBye() string
	GetNumber() int64
}

type englishBot struct {
	num int64
}
type spanishBot struct {
	num int64
}

func main() {
	eb := englishBot{7}
	sb := spanishBot{10}

	printGreeting(&eb)
	printGoodBye(&eb)
	printNumber(&eb)

	printGreeting(&sb)
	printGoodBye(&sb)
	printNumber(&sb)
}

func printGreeting(b bot) {
	fmt.Println(b.GetGreeting())
}

func printGoodBye(b bot) {
	fmt.Println(b.GetGoodBye())
}

func printNumber(b bot) {
	fmt.Println(b.GetNumber())
}

func (englishBot) GetGreeting() string {
	return "Hi there!"
}

func (englishBot) GetGoodBye() string {
	return "Good bye!"
}

func (eb *englishBot) GetNumber() int64 {
	return eb.num
}

func (spanishBot) GetGreeting() string {
	return "Hola!"
}

func (spanishBot) GetGoodBye() string {
	return "Adios!"
}

func (sb *spanishBot) GetNumber() int64 {
	return sb.num
}

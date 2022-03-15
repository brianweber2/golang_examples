package main

import (
	"log"
	"math/rand"
	"time"
)

const numPool = 1000

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

func CalculateValue(intChan chan int) {
	randNumber := RandomNumber(numPool)
	intChan <- randNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	numChannel := make(chan int)
	squaredChannel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("starting...")

	go generateNumbers(numChannel)
	go squareNumbers(numChannel, squaredChannel)
	go printNumbers(squaredChannel, &wg)

	wg.Wait()
}

func generateNumbers(numChannel chan int) {
	defer close(numChannel)
	for range 10 {
		num := rand.Intn(100)
		time.Sleep(time.Millisecond * 200)
		// fmt.Println("generated: ", num)
		numChannel <- num
	}
}

func squareNumbers(numChannel, squaredChannel chan int) {
	defer close(squaredChannel)
	for num := range numChannel {
		sq := num * num
		time.Sleep(time.Millisecond * 200)
		// fmt.Println("squared: ", sq)
		squaredChannel <- sq
	}
}

func printNumbers(squaredChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for sq := range squaredChannel {
		time.Sleep(time.Millisecond * 200)
		fmt.Println("final: ", sq)
	}
}

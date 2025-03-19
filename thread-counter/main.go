package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	rng := rand.Intn(10) + 2
	val := 0
	result := 0

	println(rng)

	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := range rng {
		wg.Add(1)
		r := rand.Intn(1000) + 100
		result += r
		fmt.Printf("go routine %d: %d\n", i, r)
		go adder(&val, &mu, &wg, r)
	}

	wg.Wait()
	println(val, result)
}

func adder(val *int, mu *sync.Mutex, wg *sync.WaitGroup, r int) {
	defer wg.Done()
	for range r {
		mu.Lock()
		*val = *val + 1
		mu.Unlock()
	}
}

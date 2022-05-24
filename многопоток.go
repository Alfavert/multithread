package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func Max1(numbers []float64) {
	defer wg.Done()
	max := 0.0
	for _, d := range numbers {
		if max < d {
			max = d
		}
	}
	fmt.Println(max)
}
func main() {
	timeStart := time.Now()
	numbers := make([]float64, 50000000)
	rand.Seed(int64(time.Now().Second()))
	for i := 0; i < 5000000; i++ {
		numbers[i] = rand.Float64() * 10
	}
	wg.Add(4)
	go Max1(numbers[0:1250000])
	go Max1(numbers[1250000:2500000])
	go Max1(numbers[2500000:3750000])
	go Max1(numbers[3750000:5000000])
	wg.Wait()
	fmt.Println(time.Since(timeStart))
}

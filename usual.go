package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		numbers []float64
		max     float64 = 0
	)

	rand.Seed(int64(time.Now().Second()))
	timeStart := time.Now()
	for i := 0; i < 500000; i++ {
		numbers[i] = rand.Float64() * 10
	}
	for d := 0; d < 500000; d++ {
		if max < numbers[d] {
			max = numbers[d]
		}
	}
	fmt.Println(max)
	fmt.Println(time.Since(timeStart))
}

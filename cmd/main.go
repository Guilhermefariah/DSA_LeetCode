package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()
	findSum(9999999999)
	TimeEnd := time.Now()

	fmt.Println(TimeEnd.Sub(timeStart).Milliseconds())
}

// Complexity
func findSum(n int) int {
	return n * (n + 1) / 2
}

// Easy
func findSumIterate(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

package main

import (
	"fmt"
	"math"
)

func count(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func count2(n int) int {
	result := 0
	for ; n > 0; n = n / 10 {
		result++
	}
	return result
}

func rescursiveCount(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + rescursiveCount(n/10)
}
func main() {
	fmt.Println(count(1234))
	fmt.Println(count2(1234))
	fmt.Println(rescursiveCount(1234))
}

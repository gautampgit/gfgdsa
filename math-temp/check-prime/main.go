package main

import (
	"fmt"
	"math"
)

func checkPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n%3 == 0 {
		return false
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkPrime(89))
}

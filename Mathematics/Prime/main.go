package main

import (
	"fmt"
	"math"
)

//naive solution time complexity is O(n)
func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//Solution with time complexity O(sqrt n)
func IsPrime2(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//effecient method for finding out whether number is prime or not
func EffecientPrime(a int) bool {
	if a == 1 {
		return false
	}
	if a == 2 || a == 3 {
		return true
	}
	if a%2 == 0 || a%3 == 0 {
		return false
	}

	for i := 5; i <= int(math.Sqrt(float64(a))); i += 6 {
		if a%i == 0 || a%(i+2) == 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println("Is Prime ", IsPrime(131))
	fmt.Println("IS Prime sqrt n time complexity", IsPrime(131))
	fmt.Println("Seive method", EffecientPrime(1032))
}

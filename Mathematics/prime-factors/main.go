package main

import (
	"fmt"
	"math"
)

func SeivePrime(a int) bool {
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

//naive solution time complexity O(n^2 log n)
func Primefactors(a int) []int {
	result := make([]int, 0)

	for i := 2; i < a; i++ {
		if SeivePrime(i) {
			x := i
			for a%x == 0 {
				result = append(result, i)
				x = x * i
			}
		}
	}
	return result
}

//solution with sqrt n itertion

func Primefactors2(a int) []int {
	result := make([]int, 0)

	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if SeivePrime(i) {
			x := i
			for a%x == 0 {
				result = append(result, i)
				x = x * i
			}
		}
	}
	return result
}

func EffecientPrime(n int) []int {
	result := make([]int, 0)

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for ; n%i == 0; n = n / i {
			result = append(result, i)
		}
	}
	if n > 1 {
		result = append(result, n)
	}
	return result
}
func moreEffecientPrime(n int) []int {
	result := make([]int, 0)
	if n <= 1 {
		return result
	}

	for ; n%2 == 0; n = n / 2 {
		result = append(result, 2)
	}
	for ; n%3 == 0; n = n / 3 {
		result = append(result, 3)
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i = i + 6 {
		for ; n%i == 0; n = n / i {
			result = append(result, i)
		}
		for ; n%(i+2) == 0; n = n / (i + 2) {
			result = append(result, (i + 2))

		}
	}
	if n > 3 {
		result = append(result, n)
	}
	return result
}

func main() {
	fmt.Println("Prime factors are ", EffecientPrime(84))
	fmt.Println("Prime Factors are", moreEffecientPrime(450))
}

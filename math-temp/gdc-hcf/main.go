package main

import (
	"fmt"
	"math"
)

//naive method to find gcd of 2 nums and time cmplexity O(min(a,b))
func findGCD(a, b int) int {
	result := int(math.Min(float64(a), float64(b)))
	for ; result > 0; result-- {
		if a%result == 0 && b%result == 0 {
			break
		}
	}
	return result
}

//eucledean method to find the gcd
func EuGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		}
		b = b - a
	}

	return a
}

//eucledean method recurssion
func EureGCD(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		return EureGCD(a-b, b)
	}
	return EureGCD(a, b-a)
}

// effecient eucledean method
func EffeEuGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return EffeEuGCD(b, a%b)
}
func main() {
	fmt.Println(EureGCD(121, 231))
	fmt.Println(EffeEuGCD(45, 90))
}

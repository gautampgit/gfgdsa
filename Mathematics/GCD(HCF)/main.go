package main

import (
	"fmt"
	"math"
)

//O(n) time complexity solution

func FindGCD(a, b int) int {
	gcd := 0
	itr := 0
	if a < b {
		itr = a
	}
	itr = b
	for i := 1; i <= itr; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}

//tutorial naive solution time complexity O(min(a,b))

func tutorialFindGCD(a, b int) int {
	result := int(math.Min(float64(a), float64(b)))
	for ; result > 0; result-- {
		if a%result == 0 && b%result == 0 {
			break
		}
	}
	return result
}

//Euclidean Algorithm using substarction
func EuclideanGCD1(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		}
		b = b - a
	}
	return a
}

//effecient euclidean method for GCD using modulous time compelxity )(log min(a,b))
func EuclideanGCD2(a, b int) int {
	if b == 0 {
		return a
	}
	return EuclideanGCD2(b, a%b)
}
func main() {
	fmt.Println("GCD ", FindGCD(320, 128))
	fmt.Println("GCD ", tutorialFindGCD(320, 128))
	//fmt.Println("GCD Euclidean minus", EuclideanGCD1(320, 128))
	fmt.Println("GCD Euclidean modulous", EuclideanGCD2(320, 128))
}

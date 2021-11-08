package main

import (
	"fmt"
	"math"
)

//naive solution time complexity O(a*b)
func LCM(a, b int) int {
	result := int(math.Max(float64(a), float64(b)))

	for true {
		if result%a == 0 && result%b == 0 {
			break
		}
		result++
	}
	return result
}

// a*b= lcm(a, b)* gcd(a, b)
// => lcm(a, b) = a*b/gcd(a*b)
//find gcd using eucledean algorithm

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func effLCM(a, b int) int {
	return (a * b) / gcd(a, b)
}
func main() {
	fmt.Println(effLCM(200, 45))
}

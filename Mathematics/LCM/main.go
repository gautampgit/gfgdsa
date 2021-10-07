package main

import (
	"fmt"
)

//naive solution time complexity O(a*b)
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func LCM(a, b int) int {
	return a * b / gcd(a, b)
}
func main() {
	fmt.Println("LCM", LCM(310, 73))
}

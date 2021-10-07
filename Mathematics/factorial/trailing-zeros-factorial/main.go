/* Counting number of trailing zeros in the factorial of a number*/
package main

import (
	"fmt"
)

//iterative solution
func trailingZers(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact = fact * i
	}
	result := 0
	for fact%10 == 0 {
		result++
		fact = fact / 10
	}
	return result
}

//counting number of 5 in the factorial of n effecient method time complexity O(n)

func trailingZeros(n int) int {
	result := 0
	for i := 5; i <= n; i = i * 5 {
		result = result + (n / i)
	}
	return result
}
func main() {
	fmt.Println("number of trailing zeros = ", trailingZeros(15))
}

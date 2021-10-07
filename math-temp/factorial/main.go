package main

import "fmt"

//recursive solution with time complexity of O(n) and space complexity O(n)
func recursiveFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

//iterative solution with time complexity of O(n) and space complexity O(1)
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = result * i
	}
	return result
}

func main() {
	fmt.Println(recursiveFactorial(4))
	fmt.Println(factorial(5))
}

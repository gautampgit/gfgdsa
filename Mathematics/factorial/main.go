/* Finding factorila of a number */
package main

import "fmt"

//recursive solution time complexity O(n) space complexity O(n)
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

//iterative solution time complexity O(n) space complexity O(1)
func iterativefactorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = result * i
	}
	return result
}
func main() {
	fmt.Println("factorial: ", factorial(6))
	fmt.Println("factorial: ", iterativefactorial(6))
}

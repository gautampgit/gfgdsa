/*Sum of 1st n natural numbers */

package main

import "fmt"

//O(1) solution
func NaturalSum(n int) int {
	return ((n / 2) * (n + 1))
}

//O(n) solution
func NaturalSum2(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
func main() {
	fmt.Println("O(1) solution", NaturalSum(6))
	fmt.Println("O(n) solution", NaturalSum2(6))
}

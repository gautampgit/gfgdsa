/*Finding number of digits in a number*/

package main

import (
	"fmt"
	"math"
)

//iterative Solution
func digits(n int) int {
	count := 0
	for n != 0 {
		n = n / 10
		count++
	}
	return count
}

//recursive Solution
func recursiveDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (1 + recursiveDigits(n/10))
}

//logarithmic solution
func logDigits(n float64) int {
	return int(math.Floor(math.Log10(n)) + 1)
}

func main() {
	fmt.Println(digits(1234))
	fmt.Println(recursiveDigits(1233))
	fmt.Println(logDigits(1234567891234567891234567))
}

package main

import "fmt"

//naive solution time complexity O( total no of bits)
func countsetbits(n int) int {
	result := 0

	for ; n > 0; n = n >> 1 {
		result = result + (n & 1)
	}
	return result
}

//brian kerningams solution time complexity O(set bits)
func briancount(n int) int {
	result := 0
	for n > 0 {
		n = n & (n - 1)
		result++
	}
	return result
}
func main() {
	fmt.Println(countsetbits(13))
	fmt.Println(briancount(5))
}

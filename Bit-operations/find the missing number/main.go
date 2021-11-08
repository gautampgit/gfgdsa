/* Given an array of size n that has values from 1 to n+1,
Hence one number in the range 1 to n+1 is missing.
 Find the missing number in the array. */

package main

import "fmt"

func FindMissing(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		result = result ^ arr[i]
	}
	for i := 0; i <= len(arr)+1; i++ {
		result = result ^ i
	}
	return result
}
func main() {
	fmt.Println(FindMissing([]int{1, 4, 3}))
}

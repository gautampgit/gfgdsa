package main

import "fmt"

//naive solution to find two odd occuring elements in an array time complexity O(n^2)
func naiveTwoOdd(arr []int) {
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count%2 != 0 {
			fmt.Println(arr[i])
		}
	}
}
func effecientTwoOdd(arr []int) {
	result, a := 0, 0
	for i := 0; i < len(arr); i++ {
		a = arr[i]
		result = result ^ arr[i]
	}
	fmt.Println(result, a)
}
func main() {
	effecientTwoOdd([]int{3, 4, 3, 4, 5, 4, 4, 6, 7, 7})
}

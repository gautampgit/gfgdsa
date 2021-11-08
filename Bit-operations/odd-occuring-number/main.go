package main

import "fmt"

//naive soltion
func OddOcuuring(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count%2 != 0 {
			result = count
		}
	}
	return result
}

//effecient solution
func EffecientOddOccuring(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		result = result ^ arr[i]
	}
	return result
}
func main() {
	fmt.Println(EffecientOddOccuring([]int{3, 4, 3, 4, 5, 5, 4, 4, 6, 7, 7}))
}

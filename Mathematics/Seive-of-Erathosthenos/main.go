package main

import "fmt"

//using the simple Seive of Erathosthenes
func simpleSeive(n int) []int {
	result := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !result[i] {
			for j := 2 * i; j <= n; j = j + i {
				result[j] = true
			}
		}
	}
	output := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !result[i] {
			output = append(output, i)
		}
	}
	return output
}
func improvedSeive(n int) []int {
	result := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !result[i] {
			for j := i * i; j <= n; j = j + i {
				result[j] = true
			}
		}
	}
	output := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !result[i] {
			output = append(output, i)
		}
	}
	return output
}
func optimizedSeive(n int) []int {
	result := make([]bool, n+1)
	output := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !result[i] {
			output = append(output, i)
			for j := i * i; j <= n; j = j + i {
				result[j] = true
			}
		}
	}
	return output
}
func main() {
	fmt.Println(simpleSeive(20))
	fmt.Println(improvedSeive(50))
	fmt.Println(optimizedSeive(100))
}

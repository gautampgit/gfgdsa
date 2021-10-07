package main

import (
	"fmt"
	"math"
)

//Naive solution O(n) tim complexity
func AllDivisors(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}

//My solution
func AllDivisors2(n int) {
	i := 1
	for ; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
	for ; i >= 1; i-- {
		if n%i == 0 {
			fmt.Println(n / i)
		}
	}

}

l04sqlche001 ZL\OINP5531660 

//solution1
func solution(n int) {
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Println(i)
			if i != n/i {
				fmt.Println(n / i)
			}
		}
	}
}
func main() {
	//AllDivisors(72)
	AllDivisors2(12)
	//solution(72)
}

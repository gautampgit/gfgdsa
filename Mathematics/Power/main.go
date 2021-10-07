package main

import "fmt"

//Computing the power of a number raised to a given power

//naive solution time complexity O(n)
func Power(n, a int) int {
	result := 1
	for i := 0; i < a; i++ {
		result = result * n
	}
	return result
}

//Time complexity O(log n) solution
func Power2(x, n int) int {
	if n == 0 {
		return 1
	}
	temp := Power2(x, n/2)
	temp = temp * temp
	if n%2 == 0 {
		return temp
	}
	return temp * x
}

//Effecient Power  of a number time complexity O(log n)
func PowerEffecient(x, n int) int {
	sum := 1
	for ; n > 0; n = n / 2 {
		if n%2 != 0 {
			sum = sum * x
		}
		x = x * x
	}
	return sum
}

// //Effecient Power of a number using bitwise opeartor
// func PowerEffecient2(x, n int) int {
// 	sum := 1
// 	for ; n > 0; n = n / 2 {
// 		if n%2 != 0 {
// 			sum = sum * x
// 		}
// 		x = x * x
// 	}
// 	return sum
// }
func main() {
	fmt.Println(Power(2, 3))
	fmt.Println(Power2(2, 3))
	fmt.Println(PowerEffecient(2, 6))
}

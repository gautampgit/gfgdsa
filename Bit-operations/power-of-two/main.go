/* check whether a number is power of 2 */

package main

import "fmt"

//naive solution
func checkPower(n int) bool {
	if n == 0 {
		return false
	}
	for ; n != 1; n = n / 2 {
		if n%2 != 0 {
			return false
		}
	}
	return true
}

//brian kerrningams solution
func checkpower(n int) bool {
	result := 0
	for n > 0 {
		n = n & (n - 1)
		result++
	}
	if result > 1 {
		return false
	}
	return true
}

//most effecient solution

func checkPowerEffecient(n int) bool {
	if n == 0 {
		return false
	}
	return (n & (n - 1)) == 0
}
func main() {
	fmt.Println(checkPower(10))
	fmt.Println(checkpower(16))
	fmt.Println(checkPowerEffecient(10))
}

/*Check wether a number is paindrome or not
a number is palindrome if its reverse is same as the number
*/

package main

import (
	"fmt"
)

func checkPalindrome(n int) bool {

	rev := 0
	rem := 0
	for temp := n; temp != 0; temp = temp / 10 {
		rem = temp % 10
		rev = rev*10 + rem

	}
	return rev == n
}
func main() {

	fmt.Println(checkPalindrome(123))
}

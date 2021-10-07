package main

import "fmt"

func Checkpalindrome(n int) bool {
	rem, rev := 0, 0
	for temp := n; temp != 0; temp = temp / 10 {
		rem = temp % 10
		rev = rev*10 + rem
	}
	return rev == n
}

func main() {
	fmt.Println(Checkpalindrome(15452))
}

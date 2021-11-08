package main

import "fmt"

func andOperation(a, b int) int {
	return a & b
}
func ORoperation(a, b int) int {
	return a | b
}
func XORoperation(a, b int) int {
	return a ^ b
}

func notOperation(a int) int {
	return ^a
}
func leftshift(a, b int) int {
	return a << b
}

func rightshift(a, b int) int {
	return a >> b
}
func main() {
	fmt.Println(andOperation(3, 6))
	fmt.Println(ORoperation(3, 6))
	fmt.Println(XORoperation(3, 6))
	fmt.Println(notOperation(5))
	fmt.Println(leftshift(-1, 3))
	fmt.Println(rightshift(-2, 1))
}

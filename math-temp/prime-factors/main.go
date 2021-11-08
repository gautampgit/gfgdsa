package main

func primeFactors(n int) []int {
	result := make([]int, 0)
	if n%2 == 0 {
		n = n / 2
		result = append(result, 2)
	}
	if n%3 == 0 {
		n = n / 3
		result = append(result, 3)
	}
	for i := 5; i <= n; i = i + 6 {
		if n%i == 0 {
			n = n / i
		}
	}
	return result
}

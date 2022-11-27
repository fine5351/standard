package main

import (
	"fmt"
)

var fibonacciResult []int

func PrintFibonacci(n int) {
	fibonacciResult = make([]int, n)
	for i := 1; i <= n; i++ {
		fmt.Printf("input: %v, output: %v\n", i, fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	if fibonacciResult[n-1] == 0 {
		fibonacciResult[n-1] = fibonacci(n-1) + fibonacci(n-2)
	}

	return fibonacciResult[n-1]
}

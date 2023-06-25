package main

import (
	"fmt"
	"strconv"
)

var fibonacci_values []int64
var prefilled_values []int64

func fibonacci(n int64) int64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func fibonacci_on_the_fly(n int64) int64 {
	if (n == 0) || (n == 1) {
		return fibonacci_values[n]
	} else {
		if fibonacci_values[n] == 0 {
			fibonacci_values[n] = fibonacci_on_the_fly(n-1) + fibonacci_on_the_fly(n-2)
			return fibonacci_values[n]
		} else {
			return fibonacci_values[n]
		}
	}
}

func initialize_slice() {
	prefilled_values = make([]int64, 93)
	prefilled_values[0] = 0
	prefilled_values[1] = 1

	for idx := 2; idx < 93; idx++ {
		prefilled_values[idx] = prefilled_values[idx-1] + prefilled_values[idx-2]
	}
}

func fibonacci_prefilled(n int64) int64 {
	return prefilled_values[n]
}

func fibonacci_bottom_up(n int64) int64 {
	if n <= 1 {
		return int64(n)
	}

	var fib_i, fib_i_minus_1, fib_i_minus_2 int64
	fib_i_minus_2 = 0
	fib_i_minus_1 = 1
	fib_i = fib_i_minus_1 + fib_i_minus_2
	for i := int64(1); i < n; i++ {
		// Calculate this value of fib_i.
		fib_i = fib_i_minus_1 + fib_i_minus_2

		// Set fib_i_minus_2 and fib_i_minus_1 for the next value.
		fib_i_minus_2 = fib_i_minus_1
		fib_i_minus_1 = fib_i
	}
	return fib_i
}

func main() {

	fibonacci_values = make([]int64, 93)
	fibonacci_values[0] = 0
	fibonacci_values[1] = 1

	initialize_slice()

	for {
		// Get n as a string.
		var n_string string
		fmt.Printf("N: ")
		fmt.Scanln(&n_string)

		// If the n string is blank, break out of the loop.
		if len(n_string) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(n_string, 10, 64)
		fmt.Printf("fibonacci_on_the_fly(%d) = %d\n", n, fibonacci_on_the_fly(n))
		//fmt.Printf("fibonacci_prefilled(%d) = %d\n", n, fibonacci_prefilled(n))
		//fmt.Printf("fibonacci_bottom_up(%d) = %d\n", n, fibonacci_bottom_up(n))
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(fibonacci_values); i++ {
		fmt.Printf("%d: %d\n", i, fibonacci_values[i])
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(prefilled_values); i++ {
		fmt.Printf("%d: %d\n", i, prefilled_values[i])
	}
}

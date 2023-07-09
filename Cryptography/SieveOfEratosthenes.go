package main

import (
	"fmt"
	"time"
)

func sieve_of_eratosthenes(max int) []bool {
	is_prime := make([]bool, max+1)
	is_prime[2] = true
	for idx := 3; idx <= max; idx += 2 {
		is_prime[idx] = true
	}

	for idx := 3; idx <= max; idx += 2 {
		if is_prime[idx] {
			for jdx := idx * 3; jdx <= max; jdx += idx {
				is_prime[jdx] = false
			}
		}
	}

	return is_prime
}

func print_sieve(sieve []bool) {
	fmt.Printf("2 ")

	for idx := 3; idx < len(sieve); idx = idx + 2 {
		if sieve[idx] {
			fmt.Printf("%d ", idx)
		}
	}
	fmt.Println()
}

func sieve_to_primes(slice []bool) []int {
	primes := []int{2}

	for idx := 3; idx < len(slice); idx = idx + 2 {
		if slice[idx] {
			primes = append(primes, idx)
		}
	}
	return primes
}

func main() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := sieve_of_eratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		print_sieve(sieve)

		primes := sieve_to_primes(sieve)
		fmt.Println(primes)
	}
}

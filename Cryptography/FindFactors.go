package main

import (
	"fmt"
	"time"
)

var primes []int

func eulers_sieve(max int) []bool {
	is_prime := make([]bool, max+1)
	is_prime[2] = true
	for idx := 3; idx <= max; idx += 2 {
		is_prime[idx] = true
	}

	for idx := 3; idx <= max; idx += 2 {
		if is_prime[idx] {
			max_q := max / idx
			if max_q%2 == 0 {
				max_q--
			}
			for qdx := max_q; qdx >= idx; qdx -= 2 {
				if is_prime[qdx] {
					is_prime[idx*qdx] = false
				}
			}
		}
	}
	return is_prime
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

func find_factors(num int) []int {

	resultMap := make(map[int]int)

	for {
		if num%2 == 0 {
			resultMap[2] = 1
			num /= 2
		} else {
			break
		}
	}

	factor := 3
	for {
		if factor*factor <= num {
			if num%factor == 0 {
				resultMap[factor] = 1
				num /= factor
			} else {
				factor += 2
			}
		} else {
			break
		}
	}

	if num > 1 {
		resultMap[num] = 1
	}

	i := 0
	result := make([]int, len(resultMap))
	for k := range resultMap {
		result[i] = k
		i++
	}

	return result
}

func multiply_slice(factors []int) int {

	ret := 1

	for _, val := range factors {
		ret *= val
	}

	return ret
}

func find_factors_sieve(num int) []int {
	var result []int

	for _, val := range primes {
		if val > num {
			break
		}
		if num%val == 0 {
			result = append(result, val)
		}
	}

	return result
}

func main() {

	var num int

	sieve := eulers_sieve(200000000)
	primes = sieve_to_primes(sieve)

	for {
		fmt.Printf("Please enter a number to prime factor: ")
		fmt.Scan(&num)

		if num < 2 {
			break
		} else {
			// Find the factors the slow way.
			start := time.Now()
			factors := find_factors_sieve(num)
			elapsed := time.Since(start)
			fmt.Printf("find_factors:       %f seconds\n", elapsed.Seconds())
			fmt.Println(factors)
			fmt.Println()
			test := multiply_slice(factors)
			fmt.Println("Factors multipled together = ", test)
		}
	}
}

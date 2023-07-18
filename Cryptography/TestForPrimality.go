package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const num_tests = 20

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func rand_range(min, max int) int {
	return min + random.Intn(max-min)
}

func fast_exp(num, pow int) int {
	result := 1
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
	}

	return result
}

func fast_exp_mod(num, pow, mod int) int {
	//return fast_exp(num, pow) % mod
	var result int = 1
	for pow > 0 {
		if pow%2 == 1 {
			result = (result * num) % mod
		}
		pow /= 2
		num = (num * num) % mod
	}
	return result
}

func is_probably_prime(num int, num_tests int) bool {
	for idx := 0; idx < num_tests; idx++ {
		n := rand_range(2, num)

		result := fast_exp_mod(n, num-1, num)

		if result != 1 {
			return false
		}
	}
	return true
}

func find_prime(min, max, num_tests int) int {
	for {
		p := rand_range(min, max+1)
		if p%2 == 0 {
			continue
		}

		if is_probably_prime(p, num_tests) {
			return p
		}
	}
}

func test_known_values() {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
		217,
	}

	probability := math.Pow(.5, num_tests)
	fmt.Printf("Probability: %f%%\n\n", probability)

	fmt.Println("Primes: ")
	for _, val := range primes {
		ret := is_probably_prime(val, num_tests)
		if ret {
			fmt.Println(val, "  Prime")
		} else {
			fmt.Println(val, "  Not Prime")
		}
	}

	fmt.Println("Composites: ")
	for _, val := range composites {
		ret := is_probably_prime(val, num_tests)
		if !ret {
			fmt.Println(val, "  Composite")
		} else {
			fmt.Println(val, "  Not Composite")
		}
	}

}

func main() {
	test_known_values()

	for {
		var digits int
		fmt.Printf("Enter a number: ")
		fmt.Scan(&digits)

		if digits == 0 {
			break
		}

		max := int(math.Pow(10, float64(digits-1)))

		prime := find_prime(1, max, num_tests)

		fmt.Println("Found Prime Number: ", prime)

	}
}

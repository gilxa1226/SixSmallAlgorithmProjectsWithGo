package main

import (
	"fmt"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func rand_range(min, max int) int {
	return min + random.Intn(max-min)
}

func lcm(a, b int) int {
	tmp := b / gcd(a, b)
	return a * tmp
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// Calculate the totient function λ(n)
// where n = p * q and p and q are prime
func totient(p, q int) int {
	return lcm(p-1, q-1)
}

// Pick a random exponent e in the range (2, λ_n)
// such that gcd(e, λ_n) = 1
func random_exponent(lambda_n int) int {
	for {
		e := rand_range(2, lambda_n)
		if gcd(e, lambda_n) == 1 {
			return e
		}
	}
	return -1
}

func inverse_mod(a, mod int) int {
	t := 0
	r := mod
	newt := 1
	newr := a

	for {
		if newr == 0 {
			break
		}
		quotient := r / newr
		t, newt = newt, t-quotient*newt
		r, newr = newr, r-quotient*newr
	}

	if r > 1 {
		panic("a is not invertible")
	}
	if t < 0 {
		t = t + mod
	}

	return t
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

func main() {
	rand.Seed(time.Now().UnixNano())

	p := find_prime(10000, 50000, 20)
	q := find_prime(10000, 50000, 20)

	n := p * q
	t := totient(p, q)
	e := random_exponent(t)
	d := inverse_mod(e, t)

	fmt.Printf("\n\n")
	fmt.Println("*** Public ***")
	fmt.Println("Public key modulus: ", n)
	fmt.Println("Public key exponent e: ", e)

	fmt.Printf("\n*** Private ***\n")
	fmt.Printf("Primes:\t%d, %d\n", p, q)
	fmt.Println("Lambda(n):  ", t)
	fmt.Println("d: ", d)

	var m int

	for {
		fmt.Printf("Message: ")
		fmt.Scan(&m)

		if m < 1 || m > n-1 {
			break
		}

		encrypted := fast_exp_mod(m, e, n)
		fmt.Println("Ciphertext: ", encrypted)

		decrypted := fast_exp_mod(encrypted, d, n)
		fmt.Println("Decrypted: ", decrypted)
	}
}

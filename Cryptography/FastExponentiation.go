package main

import (
	"fmt"
	"math"
)

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
	return fast_exp(num, pow) % mod
}

func main() {

	var num, pow, mod int

	for {
		num, pow, mod = -1, -1, -1
		fmt.Printf("Enter Num: ")
		fmt.Scanln(&num)
		fmt.Printf("Enter Pow: ")
		fmt.Scanln(&pow)
		fmt.Printf("Enter Mod: ")
		fmt.Scanln(&mod)
		if num < 0 || pow < 0 || mod < 0 {
			break
		}

		fmt.Printf("Fast Exp: %d\n", fast_exp(num, pow))
		fmt.Printf("math.Pow: %f\n", math.Pow(float64(num), float64(pow)))

		fmt.Printf("Fast Exp Mod: %d\n", fast_exp_mod(num, pow, mod))
		fmt.Printf("math.Pow Mod: %d\n", int(math.Pow(float64(num), float64(pow)))%mod)
	}

}

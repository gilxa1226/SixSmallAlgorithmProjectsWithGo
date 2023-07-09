package main

import (
	"fmt"
	"strconv"
)

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

func lcm(a, b int) int {
	tmp := b / gcd(a, b)
	return a * tmp
}

func main() {

	var A string
	var B string

	for {
		A = ""
		B = ""
		fmt.Printf("Enter A: ")
		fmt.Scanln(&A)
		fmt.Printf("Enter B: ")
		fmt.Scanln(&B)
		if A == "" || B == "" {
			break
		}

		intA, _ := strconv.Atoi(A)
		intB, _ := strconv.Atoi(B)

		fmt.Println("A\tB\tGCD(A,B)\tLCM(A,B)")
		fmt.Printf("%d\t%d\t%d\t%d\n", intA, intB, gcd(intA, intB), lcm(intA, intB))
	}

}

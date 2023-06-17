package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func make_random_array(num_items, max int) []int {
	var array = make([]int, num_items)

	for idx := 0; idx < num_items; idx++ {
		array[idx] = rand.Intn(max)
	}

	return array
}

func print_array(arr []int, num_items int) {
	if len(arr) == num_items {
		fmt.Println("Printing full array")
		fmt.Println(arr)
	} else {
		fmt.Printf("Printing first %d elements of the array \n", num_items)
		for idx := 0; idx < num_items; idx++ {
			fmt.Printf("%d ", arr[idx])
		}
		fmt.Println()
	}
}

func linear_search(arr []int, target int) (index, num_tests int) {
	for idx, num := range arr {
		if num == target {
			return idx, idx
		}
	}

	return -1, len(arr)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var num_items, max int
	var toFind string
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted array.
	arr := make_random_array(num_items, max)
	print_array(arr, 40)

	for {
		fmt.Printf("Enter a number to search for: ")
		fmt.Scanln(&toFind)
		if toFind == "" {
			break
		}

		iToFind, _ := strconv.Atoi(toFind)
		idx, tests := linear_search(arr, iToFind)

		if idx == -1 {
			fmt.Printf("Target: %d not found, %d tests\n", iToFind, tests)
		} else {
			fmt.Printf("Target: %d found at index: %d\n", iToFind, idx)
		}
	}

}

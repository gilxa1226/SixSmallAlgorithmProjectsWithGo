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

func partition(arr []int) int {
	lo := 0
	hi := len(arr) - 1
	j := lo

	pivot := arr[hi]

	i := lo - 1

	for j = lo; j <= hi-1; j++ {
		if arr[j] <= pivot {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	i = i + 1
	arr[i], arr[j] = arr[j], arr[i]
	return i
}

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	p := partition(arr)

	quicksort(arr[:p])
	quicksort(arr[p:])
}

func binary_search(arr []int, target int) (index, num_tests int) {
	min := 0
	max := len(arr) - 1
	tests := 0

	for {
		tests++
		avg := (min + max) / 2

		if min >= max {
			break
		}

		if arr[avg] == target {
			return avg, tests
		} else if arr[avg] > target {
			max = avg - 1
		} else if arr[avg] < target {
			min = avg + 1
		}
	}

	return -1, tests
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
	quicksort(arr)
	print_array(arr, 40)

	for {
		toFind = ""
		fmt.Printf("Enter a number to search for: ")
		fmt.Scanln(&toFind)
		if toFind == "" {
			break
		}

		iToFind, _ := strconv.Atoi(toFind)
		idx, tests := binary_search(arr, iToFind)

		if idx == -1 {
			fmt.Printf("Target: %d not found, %d tests\n", iToFind, tests)
		} else {
			fmt.Printf("Target: %d found at index: %d in %d tests\n", iToFind, idx, tests)
		}
	}

}

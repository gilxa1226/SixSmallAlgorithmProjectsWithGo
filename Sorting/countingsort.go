package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Customer struct {
    id              string
    num_purchases   int
}

func make_random_array(num_items, max int) []Customer {
	var array = make([]Customer, num_items)

	for idx := 0; idx < num_items; idx++ {
	    array[idx] = Customer{ id: "C"+strconv.Itoa(idx), num_purchases: rand.Intn(max)}
	}

	return array
}

func print_array(arr []Customer, num_items int) {
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

func check_sorted(arr []Customer) {
	length := len(arr)

	for idx := 1; idx < length; idx++ {
		if arr[idx-1].num_purchases > arr[idx].num_purchases {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}

	fmt.Println("The array is sorted")
}

func zero_array(arr []int) {
	for i := range arr {
		arr[i] = 0
	}
}

func countingsort(arr []Customer, max int) []Customer {
	var counts = make([]int, max+1)
	zero_array(counts)

	for _, val := range arr {
		counts[val.num_purchases]++
	}

	for idx := 1; idx < max; idx++ {
		counts[idx] += counts[idx-1]
	}

	var sorted = make([]Customer, len(arr))

	for idx := len(arr) - 1; idx >= 0; idx-- {
		aval := arr[idx]         // a[i]
		cval := counts[aval.num_purchases] - 1 // [c[a[i]]
		sorted[cval] = aval      // b[c[a[i]]] = a[i]
		counts[aval.num_purchases]--
	}

	return sorted

}

func main() {

	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted array.
	arr := make_random_array(num_items, max)
	print_array(arr, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingsort(arr, max)
	print_array(sorted, 40)

	// Verify that it's sorted.
	check_sorted(sorted)
}

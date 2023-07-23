package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const num_items = 340 // A reasonable value for exhaustive search.

const min_value = 1
const max_value = 10
const min_weight = 4
const max_weight = 10

var allowed_weight int

type Item struct {
	id, blocked_by int
	block_list     []int
	value, weight  int
	is_selected    bool
}

func make_items(num_items, min_value, max_value, min_weight, max_weight int) []Item {
	// Initialize a pseudorandom number generator
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Changing seed
	// random := rand.New(rand.NewSource(1337)) // fixed seed for testing

	items := make([]Item, num_items)
	for i := 0; i < num_items; i++ {
		items[i] = Item{
			i,
			-1,
			nil,
			random.Intn(max_value-min_value+1) + min_value,
			random.Intn(max_weight-min_weight+1) + min_weight,
			false,
		}
	}
	return items
}

// Return a copy of the items slice
func copy_items(items []Item) []Item {
	new_items := make([]Item, len(items))
	copy(new_items, items)
	return new_items
}

// Return the total value of the items.
// If add_all is false, only add up the selected items.
func sum_values(items []Item, add_all bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if add_all || items[i].is_selected {
			total += items[i].value
		}
	}
	return total
}

// Return the total weight of the items.
// If add_all is false, only add up the selected items.
func sum_weights(items []Item, add_all bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if add_all || items[i].is_selected {
			total += items[i].weight
		}
	}
	return total
}

// Return the value of this solution.
// If the solution is too heavy, return -1 so we prefer an empty solution.
func solution_value(items []Item, allowed_weight int) int {
	// If the solution's total weight > allowed_weight,
	// return -1 so we won't use this solution.
	if sum_weights(items, false) > allowed_weight {
		return -1
	}

	// Return the sum of the selected values.
	return sum_values(items, false)
}

// Print the selected items.
func print_selected(items []Item) {
	num_printed := 0
	for i, item := range items {
		if item.is_selected {
			fmt.Printf("%d(%d, %d) ", i, item.value, item.weight)
		}
		num_printed += 1
		if num_printed > 100 {
			fmt.Println("...")
			return
		}
	}
	fmt.Println()
}

func run_algorithm(alg func([]Item, int) ([]Item, int, int), items []Item, allowed_weight int) {
	// Copy the items so the run isn't influenced by a previous run.
	test_items := copy_items(items)

	start := time.Now()

	// Run the algorithm.
	solution, total_value, function_calls := alg(test_items, allowed_weight)

	elapsed := time.Since(start)

	fmt.Printf("Elapsed: %f\n", elapsed.Seconds())
	print_selected(solution)
	fmt.Printf("Value: %d, Weight: %d, Calls: %d\n",
		total_value, sum_weights(solution, false), function_calls)
	fmt.Println()
}

// Recursively assign values in or out of the solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func branch_and_bound(items []Item, allowed_weight int) ([]Item, int, int) {
	best_value := 0
	current_value := 0
	current_weight := 0
	remaining_value := sum_values(items, true)

	return do_branch_and_bound(items, allowed_weight, 0,
		best_value, current_value, current_weight, remaining_value)
}

func do_branch_and_bound(items []Item, allowed_weight, next_index, best_value, current_value,
	current_weight, remaining_value int) ([]Item, int, int) {

	var test1_solution, test2_solution []Item
	var test1_value, test1_calls, test2_value, test2_calls int

	if next_index >= len(items) {
		return items, current_value, 1
	}

	if current_value+remaining_value <= best_value {
		return nil, 0, 1
	}

	if current_value+items[next_index].weight <= allowed_weight {
		items[next_index].is_selected = true
		test1_solution, test1_value, test1_calls = do_branch_and_bound(items, allowed_weight, next_index+1, best_value,
			current_value+items[next_index].value, current_weight+items[next_index].weight,
			remaining_value-items[next_index].value)
	} else {
		test1_solution = nil
		test1_value = 0
		test1_calls = 1
	}

	items[next_index].is_selected = false
	test2_solution, test2_value, test2_calls = do_branch_and_bound(items, allowed_weight, next_index+1,
		best_value, current_value, current_weight, remaining_value)

	//if solution_value(test1_solution, allowed_weight) > solution_value(test2_solution, allowed_weight) {
	if test1_value > test2_value {
		items[next_index].is_selected = true
		return test1_solution, test1_value, test1_calls + 1
	} else {
		return test2_solution, test2_value, test2_calls + 1
	}
}

func rods_technique(items []Item, allowed_weight int) ([]Item, int, int) {
	best_value := 0
	current_value := 0
	current_weight := 0
	remaining_value := sum_values(items, true)

	make_block_lists(items)

	return do_rods_technique(items, allowed_weight, 0,
		best_value, current_value, current_weight, remaining_value)
}

func rods_technique_sorted(items []Item, allowed_weight int) ([]Item, int, int) {
	best_value := 0
	current_value := 0
	current_weight := 0
	remaining_value := sum_values(items, true)

	make_block_lists(items)

	sort.Slice(items, func(i, j int) bool {
		return len(items[i].block_list) > len(items[j].block_list)
	})

	for idx := range items {
		items[idx].id = idx
	}

	make_block_lists(items)

	return do_rods_technique(items, allowed_weight, 0,
		best_value, current_value, current_weight, remaining_value)
}

func do_rods_technique(items []Item, allowed_weight, next_index, best_value, current_value,
	current_weight, remaining_value int) ([]Item, int, int) {

	var test1_solution, test2_solution []Item
	var test1_value, test1_calls, test2_value, test2_calls int

	if next_index >= len(items) {
		return items, current_value, 1
	}

	if current_value+remaining_value <= best_value {
		return nil, 0, 1
	}

	if current_value+items[next_index].weight <= allowed_weight {
		items[next_index].is_selected = true
		test1_solution = nil
		test1_value = 0
		test1_calls = 1
		if items[next_index].blocked_by < 0 {
			test1_solution, test1_value, test1_calls = do_rods_technique(items, allowed_weight, next_index+1, best_value,
				current_value+items[next_index].value, current_weight+items[next_index].weight,
				remaining_value-items[next_index].value)
		}
	} else {
		test1_solution = nil
		test1_value = 0
		test1_calls = 1
	}

	block_items(items[next_index], items)
	items[next_index].is_selected = false
	test2_solution, test2_value, test2_calls = do_rods_technique(items, allowed_weight, next_index+1,
		best_value, current_value, current_weight, remaining_value)
	unblock_items(items[next_index], items)

	//if solution_value(test1_solution, allowed_weight) > solution_value(test2_solution, allowed_weight) {
	if test1_value > test2_value {
		items[next_index].is_selected = true
		return test1_solution, test1_value, test1_calls + 1
	} else {
		return test2_solution, test2_value, test2_calls + 1
	}
}

// build the item's block lists
func make_block_lists(items []Item) {
	for i, ival := range items {
		items[i].block_list = make([]int, 0)
		for j, jval := range items {
			if i != j && ival.value >= jval.value && ival.weight <= jval.weight {
				items[i].block_list = append(items[i].block_list, j)
			}
		}
	}
}

// Block items on this item's blocks list.
func block_items(source Item, items []Item) {
	for _, val := range source.block_list {
		if items[val].blocked_by < 0 {
			items[val].blocked_by = source.id
		}
	}
}

// Unblock items on the item's blocks list
func unblock_items(source Item, items []Item) {
	for _, val := range source.block_list {
		if items[val].blocked_by == source.id {
			items[val].blocked_by = -1
		}
	}
}

func main() {
	items := make_items(num_items, min_value, max_value, min_weight, max_weight)
	allowed_weight = sum_weights(items, true) / 2

	// Display basic parameters.
	fmt.Println("*** Parameters ***")
	fmt.Printf("# items: %d\n", num_items)
	fmt.Printf("Total value: %d\n", sum_values(items, true))
	fmt.Printf("Total weight: %d\n", sum_weights(items, true))
	fmt.Printf("Allowed weight: %d\n", allowed_weight)
	fmt.Println()

	// Exhaustive search
	if num_items > 25 { // Only run exhaustive search if num_items <= 25.
		fmt.Println("Too many items for exhaustive search\n")
	} else {
		fmt.Println("*** Exhaustive Search ***")
		run_algorithm(branch_and_bound, items, allowed_weight)
	}

	// Rod's technique
	if num_items > 85 { // Only use Rod's technique if num_items <= 85.
		fmt.Println("Too many items for Rod's technique\n")
	} else {
		fmt.Println("*** Rod's technique ***")
		run_algorithm(rods_technique, items, allowed_weight)
	}

	// Rod's technique sorted
	if num_items > 350 { // Only use Rod's technique if num_items <= 85.
		fmt.Println("Too many items for Rod's technique sorted\n")
	} else {
		fmt.Println("*** Rod's technique sorted ***")
		run_algorithm(rods_technique, items, allowed_weight)
	}
}

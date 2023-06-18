package main

import "fmt"

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	me.next = after
}

func (me *Cell) delete_after() *Cell {
	if me.next == nil {
		panic("No cell to delete after me")
	}

	delete := me.next
	me.next = delete.next

	return delete
}

func (list *LinkedList) add_range(values []string) {

	var last_cell *Cell

	if list.sentinel.next == nil {
		last_cell = list.sentinel
	} else {
		last_cell = list.sentinel.next
	}

	for {
		if last_cell.next == nil {
			break
		} else {
			last_cell = last_cell.next
		}
	}

	for _, str := range values {
		tmp := Cell{str, nil}
		last_cell.add_after(&tmp)
		last_cell = last_cell.next
	}
}

func (list *LinkedList) to_string(separator string) string {

	var last_cell *Cell

	if list.sentinel.next == nil {
		last_cell = list.sentinel
	} else {
		last_cell = list.sentinel.next
	}
	retString := ""

	for {
		if last_cell.next == nil {
			retString += last_cell.data
			return retString
		} else {
			retString += last_cell.data + separator
			last_cell = last_cell.next
		}
	}
}

func (list *LinkedList) length() int {

	count := 0
	last_cell := list.sentinel

	for {
		if last_cell.next == nil {
			return count
		} else {
			count++
			last_cell = last_cell.next
		}
	}
}

func (list *LinkedList) is_empty() bool {
	if list.sentinel.next == nil {
		return true
	} else {
		return false
	}
}

func (list *LinkedList) push(data string) {
	tmp := Cell{data, list.sentinel.next}
	list.sentinel.next = &tmp
}

func (list *LinkedList) pop() string {
	str := list.sentinel.next.data
	list.sentinel = list.sentinel.next
	return str
}

func make_linked_list() *LinkedList {
	tmp := LinkedList{sentinel: &Cell{data: "", next: nil}}
	return &tmp
}

func main() {
	// small_list_test()

	// Make a list from an array of values.
	greek_letters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := make_linked_list()
	list.add_range(greek_letters)
	fmt.Println(list.to_string(" "))
	fmt.Println()

	// Demonstrate a stack.
	stack := make_linked_list()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.is_empty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.to_string(" "))
	}
}

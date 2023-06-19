package main

import "fmt"

type Cell struct {
	data     string
	forward  *Cell
	backward *Cell
}

type DoublyLinkedList struct {
	top_sentinal    *Cell
	bottom_sentinal *Cell
}

func (me *Cell) add_after(after *Cell) {
	after.forward = me.forward
	after.backward = me
	me.forward = after
}

func (me *Cell) delete_after() *Cell {
	if me.forward == nil {
		panic("No cell to delete after me")
	}

	delete := me.forward
	me.forward = delete.forward
	delete.forward.backward = me

	return delete
}

func (list *DoublyLinkedList) add_range(values []string) {

	var last_cell *Cell

	if list.top_sentinal == list.bottom_sentinal {
		// list is empty
		last_cell = list.top_sentinal
	} else {
		last_cell = list.bottom_sentinal.backward
	}

	for _, str := range values {
		tmp := Cell{str, nil, nil}
		last_cell.add_after(&tmp)
		last_cell = last_cell.forward
	}

	list.bottom_sentinal.backward = last_cell

}

func (list *DoublyLinkedList) to_string(separator string) string {

	retString := ""

	if list.is_empty() {
		return ""
	}

	last_cell := list.top_sentinal.forward

	for {

		if last_cell.forward == nil {
			retString += last_cell.data
			return retString
		} else {
			retString += last_cell.data + separator
			last_cell = last_cell.forward
		}
	}
}

func (list *DoublyLinkedList) to_string_max(separator string, max int) string {

	retString := ""
	count := 0

	if list.is_empty() {
		return ""
	}

	last_cell := list.top_sentinal.forward

	for {
		if count == max {
			return retString
		}
		if last_cell.forward == nil {
			retString += last_cell.data
			return retString
		} else {
			retString += last_cell.data + separator
			last_cell = last_cell.forward
		}
		count++
	}
}

func (list *DoublyLinkedList) length() int {

	if list.is_empty() {
		return 0
	}

	count := 0
	last_cell := list.top_sentinal.forward

	for {
		if last_cell.forward == nil {
			return count
		} else {
			count++
			last_cell = last_cell.forward
		}
	}

	return count
}

func (list *DoublyLinkedList) is_empty() bool {
	if list.top_sentinal.forward == list.bottom_sentinal {
		return true
	} else {
		return false
	}
}

func (list *DoublyLinkedList) push(data string) {
	curtop := list.top_sentinal.forward
	tmp := Cell{data, curtop, list.top_sentinal}
	curtop.backward = &tmp
	list.top_sentinal.forward = &tmp
}

func (list *DoublyLinkedList) pop() string {
	str := list.top_sentinal.forward.data
	list.top_sentinal = list.top_sentinal.forward
	list.top_sentinal.forward.backward = list.top_sentinal
	return str
}

func (list *DoublyLinkedList) enqueue(data string) {
	list.push(data)
}

func (list *DoublyLinkedList) dequeue() string {
	if list.is_empty() {
		panic("Queue is empty")
	}

	ret := list.bottom_sentinal.backward.data
	list.bottom_sentinal.backward = list.bottom_sentinal.backward.backward
	list.bottom_sentinal.backward.forward = list.bottom_sentinal

	return ret
}

func (list *DoublyLinkedList) push_bottom(data string) {
	tmp := Cell{data: data, forward: list.bottom_sentinal, backward: list.bottom_sentinal.backward}
	list.bottom_sentinal.backward.forward = &tmp
	list.bottom_sentinal.backward = &tmp
}

func (list *DoublyLinkedList) push_top(data string) {
	list.push(data)
}

func (list *DoublyLinkedList) pop_top() string {
	return list.pop()
}

func (list *DoublyLinkedList) pop_bottom() string {
	return list.dequeue()
}

func make_doubly_linked_list() *DoublyLinkedList {
	top_sentinal := Cell{data: "", forward: nil, backward: nil}
	bottom_sentinal := Cell{data: "", forward: nil, backward: nil}
	top_sentinal.forward = &bottom_sentinal
	bottom_sentinal.backward = &top_sentinal
	tmp := DoublyLinkedList{top_sentinal: &top_sentinal, bottom_sentinal: &bottom_sentinal}
	return &tmp
}

func main() {
	// Make a list from a slice of values.
	list := make_doubly_linked_list()
	animals := []string{
		"Ant",
		"Bat",
		"Cat",
		"Dog",
		"Elk",
		"Fox",
	}
	list.add_range(animals)
	fmt.Println(list.to_string(" "))

	// Test queue functions.
	fmt.Printf("*** Queue Functions ***\n")
	queue := make_doubly_linked_list()
	queue.enqueue("Agate")
	queue.enqueue("Beryl")
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Citrine")
	fmt.Printf("%s ", queue.dequeue())
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Diamond")
	queue.enqueue("Emerald")
	for !queue.is_empty() {
		fmt.Printf("%s ", queue.dequeue())
	}
	fmt.Printf("\n\n")

	// Test deque functions. Names starting
	// with F have a fast pass.
	fmt.Printf("*** Deque Functions ***\n")
	deque := make_doubly_linked_list()
	deque.push_top("Ann")
	deque.push_top("Ben")
	fmt.Printf("%s ", deque.pop_bottom())
	deque.push_bottom("F-Cat")
	fmt.Printf("%s ", deque.pop_bottom())
	fmt.Printf("%s ", deque.pop_bottom())
	deque.push_bottom("F-Dan")
	deque.push_top("Eva")
	for !deque.is_empty() {
		fmt.Printf("%s ", deque.pop_bottom())
	}
	fmt.Printf("\n")
}

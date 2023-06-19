package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

type Cell struct {
	data     *Node
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

func (list *DoublyLinkedList) push(data Node) {
	curtop := list.top_sentinal.forward
	tmp := Cell{&data, curtop, list.top_sentinal}
	curtop.backward = &tmp
	list.top_sentinal.forward = &tmp
}

func (list *DoublyLinkedList) pop() Node {
	str := list.top_sentinal.forward.data
	list.top_sentinal = list.top_sentinal.forward
	list.top_sentinal.forward.backward = list.top_sentinal
	return *str
}

func (list *DoublyLinkedList) enqueue(data Node) {
	list.push(data)
}

func (list *DoublyLinkedList) dequeue() Node {
	if list.is_empty() {
		panic("Queue is empty")
	}

	ret := list.bottom_sentinal.backward.data
	list.bottom_sentinal.backward = list.bottom_sentinal.backward.backward
	list.bottom_sentinal.backward.forward = list.bottom_sentinal

	return *ret
}

func (list *DoublyLinkedList) push_bottom(data Node) {
	tmp := Cell{data: &data, forward: list.bottom_sentinal, backward: list.bottom_sentinal.backward}
	list.bottom_sentinal.backward.forward = &tmp
	list.bottom_sentinal.backward = &tmp
}

func (list *DoublyLinkedList) push_top(data Node) {
	list.push(data)
}

func (list *DoublyLinkedList) pop_top() Node {
	return list.pop()
}

func (list *DoublyLinkedList) pop_bottom() Node {
	return list.dequeue()
}

func make_doubly_linked_list() *DoublyLinkedList {
	top_sentinal := Cell{data: &Node{"", nil, nil}, forward: nil, backward: nil}
	bottom_sentinal := Cell{data: &Node{"", nil, nil}, forward: nil, backward: nil}
	top_sentinal.forward = &bottom_sentinal
	bottom_sentinal.backward = &top_sentinal
	tmp := DoublyLinkedList{top_sentinal: &top_sentinal, bottom_sentinal: &bottom_sentinal}
	return &tmp
}

func (node *Node) display_indented(indent string, depth int) string {
	result := strings.Repeat(indent, depth) + node.data + "\n"

	if node.left != nil {
		result = result + node.left.display_indented(indent, depth+1)
	}
	if node.right != nil {
		result = result + node.right.display_indented(indent, depth+1)
	}

	return result
}

func (node *Node) preorder() string {
	result := node.data

	if node.left != nil {
		result = result + " " + node.left.preorder()
	}
	if node.right != nil {
		result = result + " " + node.right.preorder()
	}

	return result
}

func (node *Node) inorder() string {
	result := ""

	if node.left != nil {
		result = result + node.left.inorder()
	}

	result = result + " " + node.data

	if node.right != nil {
		result = result + node.right.inorder()
	}

	return result
}

func (node *Node) postorder() string {
	result := ""

	if node.left != nil {
		result = result + node.left.postorder()
	}

	if node.right != nil {
		result = result + node.right.postorder()
	}

	result = result + " " + node.data

	return result
}

func (node *Node) breadth_first() string {

	queue := make_doubly_linked_list()
	result := ""
	queue.enqueue(*node)

	for {
		if queue.is_empty() {
			return result
		}
		tmp := queue.dequeue()
		result = result + tmp.data + " "
		if tmp.left != nil {
			queue.enqueue(*tmp.left)
		}
		if tmp.right != nil {
			queue.enqueue(*tmp.right)
		}
	}

}

func build_tree() *Node {
	a_node := Node{"A", nil, nil}
	b_node := Node{"B", nil, nil}
	c_node := Node{"C", nil, nil}
	d_node := Node{"D", nil, nil}
	e_node := Node{"E", nil, nil}
	f_node := Node{"F", nil, nil}
	g_node := Node{"G", nil, nil}
	h_node := Node{"H", nil, nil}
	i_node := Node{"I", nil, nil}
	j_node := Node{"J", nil, nil}
	a_node.left = &b_node
	b_node.left = &d_node
	b_node.right = &e_node
	e_node.left = &g_node

	a_node.right = &c_node
	c_node.right = &f_node
	f_node.left = &h_node
	h_node.left = &i_node
	h_node.right = &j_node
	return &a_node
}

func main() {
	tree := build_tree()
	fmt.Println(tree.display_indented("  ", 0))
	fmt.Println("Preorder:   ", tree.preorder())
	fmt.Println("Preorder:   ", tree.inorder())
	fmt.Println("Preorder:   ", tree.postorder())
	fmt.Println("Preorder:   ", tree.breadth_first())
}

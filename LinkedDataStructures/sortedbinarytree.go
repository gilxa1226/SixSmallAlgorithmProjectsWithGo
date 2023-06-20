package main

import "fmt"

type Node struct {
	data  string
	left  *Node
	right *Node
}

func (node *Node) insert_value(data string) {

	for {
		if data <= node.data {
			// Go left
			if node.left != nil {
				node.left.insert_value(data)
			} else {
				new_node := Node{data, nil, nil}
				node.left = &new_node
				break
			}
		} else {
			// Go right
			if node.right != nil {
				node.right.insert_value(data)
			} else {
				new_node := Node{data, nil, nil}
				node.right = &new_node
				break
			}
		}
		break
	}

	return
}

func (node *Node) find_value(data string) *Node {
	for {
		if data == node.data {
			return node
		} else if data < node.data {
			// Go left
			if node.left != nil {
				return node.left.find_value(data)
			} else {
				return nil
			}
		} else {
			// Go right
			if node.right != nil {
				return node.right.find_value(data)
			} else {
				return nil
			}
		}
	}
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

func main() {
	// Make a root node to act as sentinel.
	root := Node{"", nil, nil}

	// Add some values.
	root.insert_value("I")
	root.insert_value("G")
	root.insert_value("C")
	root.insert_value("E")
	root.insert_value("B")
	root.insert_value("K")
	root.insert_value("S")
	root.insert_value("Q")
	root.insert_value("M")

	// Add F.
	root.insert_value("F")

	// Display the values in sorted order.
	fmt.Printf("Sorted values: %s\n", root.right.inorder())

	// Let the user search for values.
	for {
		// Get the target value.
		target := ""
		fmt.Printf("String: ")
		fmt.Scanln(&target)
		if len(target) == 0 {
			break
		}

		// Find the value's node.
		node := root.find_value(target)
		if node == nil {
			fmt.Printf("%s not found\n", target)
		} else {
			fmt.Printf("Found value %s\n", target)
		}
	}

}

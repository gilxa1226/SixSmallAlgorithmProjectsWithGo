package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Employee struct {
	name  string
	phone string
}

type LinearProbingHashTable struct {
	capacity  int
	employees []*Employee
}

// djb2 hash function. See http://www.sce.yorku.ca/~oz/hash.html
func hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func NewLinearProbingHashTable(capacity int) *LinearProbingHashTable {
	empArr := make([]*Employee, capacity)
	return &LinearProbingHashTable{capacity: capacity, employees: empArr}
}

func (hashTable *LinearProbingHashTable) dump() {
	for idx := 0; idx < hashTable.capacity; idx++ {
		if hashTable.employees[idx] != nil {
			fmt.Printf("%2d: %s\t%s\n", idx, hashTable.employees[idx].name, hashTable.employees[idx].phone)
		} else {
			fmt.Printf("%2d: ---\n", idx)
		}
	}
}

func (hashTable *LinearProbingHashTable) find(name string) (int, int) {
	hash := hash(name) % hashTable.capacity

	for i := 0; i < hashTable.capacity; i++ {
		idx := (hash + i) % hashTable.capacity
		if hashTable.employees[idx] == nil || hashTable.employees[idx].name == name {
			return idx, i
		}
	}

	return -1, -1
}

func (hashTable *LinearProbingHashTable) set(name string, phone string) {
	idx, _ := hashTable.find(name)

	if idx < 0 {
		panic("Key is not in the table")
	}

	if hashTable.employees[idx] != nil {
		hashTable.employees[idx].phone = phone
	} else {
		emp := Employee{name: name, phone: phone}
		hashTable.employees[idx] = &emp
	}
}

func (hashTable *LinearProbingHashTable) get(name string) string {
	idx, _ := hashTable.find(name)

	if idx < 0 {
		return ""
	}

	if hashTable.employees[idx] == nil {
		return ""
	}

	return hashTable.employees[idx].phone
}

func (hashTable *LinearProbingHashTable) contains(name string) bool {
	idx, _ := hashTable.find(name)

	if idx < 0 || hashTable.employees[idx] == nil {
		return false
	}

	return true

}

// Make a display showing whether each slice entry is nil.
func (hashTable *LinearProbingHashTable) dump_concise() {
	// Loop through the slice.
	for i, employee := range hashTable.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else {
			// Display this entry.
			fmt.Printf("O")
		}
		if i%50 == 49 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Return the average probe sequence length for the items in the table.
func (hashTable *LinearProbingHashTable) ave_probe_sequence_length() float32 {
	total_length := 0
	num_values := 0
	for _, employee := range hashTable.employees {
		if employee != nil {
			_, probe_length := hashTable.find(employee.name)
			total_length += probe_length
			num_values++
		}
	}
	return float32(total_length) / float32(num_values)
}

func main() {
	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101"},
		Employee{"Bob Baker", "202-555-0102"},
		Employee{"Cindy Cant", "202-555-0103"},
		Employee{"Dan Deever", "202-555-0104"},
		Employee{"Edwina Eager", "202-555-0105"},
		Employee{"Fred Franklin", "202-555-0106"},
		Employee{"Gina Gable", "202-555-0107"},
	}

	hash_table := NewLinearProbingHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	// fmt.Println("Deleting Dan Deever")
	// hash_table.delete("Dan Deever")
	// fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))

	// Look at clustering.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	big_capacity := 1009
	big_hash_table := NewLinearProbingHashTable(big_capacity)
	num_items := int(float32(big_capacity) * 0.9)
	for i := 0; i < num_items; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		big_hash_table.set(str, str)
	}
	big_hash_table.dump_concise()
	fmt.Printf("Average probe sequence length: %f\n",
		big_hash_table.ave_probe_sequence_length())

}

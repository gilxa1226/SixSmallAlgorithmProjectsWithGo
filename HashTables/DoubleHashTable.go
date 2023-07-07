package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Employee struct {
	name    string
	phone   string
	deleted bool
}

type DoubleHashTable struct {
	capacity  int
	employees []*Employee
}

// djb2 hash function. See http://www.sce.yorku.ca/~oz/hash.html
func hash1(value string) int {
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

func hash2(value string) int {
	hash := 0
	for _, ch := range value {
		hash += int(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}

	// Make sure the result is not 0.
	if hash == 0 {
		hash = 1
	}
	return hash
}

func NewDoubleHashTable(capacity int) *DoubleHashTable {
	empArr := make([]*Employee, capacity)
	return &DoubleHashTable{capacity: capacity, employees: empArr}
}

func (hashTable *DoubleHashTable) dump() {
	for idx := 0; idx < hashTable.capacity; idx++ {
		if hashTable.employees[idx] != nil {
			if !hashTable.employees[idx].deleted {
				fmt.Printf("%2d: %s\t%s\n", idx, hashTable.employees[idx].name, hashTable.employees[idx].phone)
			} else {
				fmt.Printf("%2d: xxx\n", idx)
			}
		} else {
			fmt.Printf("%2d: ---\n", idx)
		}
	}
}

func (hashTable *DoubleHashTable) find(name string) (int, int) {
	hash1 := hash2(name) % hashTable.capacity
	hash2 := hash2(name) % hashTable.capacity
	deleted_index := -1

	for i := 0; i < hashTable.capacity; i++ {
		idx := (hash1 + i + hash2) % hashTable.capacity
		if hashTable.employees[idx] == nil {
			if deleted_index >= 0 {
				return deleted_index, i
			} else {
				return idx, i
			}
		} else if deleted_index == -1 && hashTable.employees[idx].deleted {
			deleted_index = idx
		} else if hashTable.employees[idx].name == name {
			return idx, i
		}
	}

	if deleted_index >= 0 {
		return deleted_index, -1
	}

	return -1, -1
}

func (hashTable *DoubleHashTable) set(name string, phone string) {
	idx, _ := hashTable.find(name)

	if idx < 0 {
		panic("Key is not in the table, and no deleted items")
	}

	if hashTable.employees[idx] != nil {
		hashTable.employees[idx].phone = phone
	} else {
		emp := Employee{name: name, phone: phone, deleted: false}
		hashTable.employees[idx] = &emp
	}
}

func (hashTable *DoubleHashTable) get(name string) string {
	idx, _ := hashTable.find(name)

	if idx < 0 {
		return ""
	}

	if hashTable.employees[idx] == nil {
		return ""
	}

	if hashTable.employees[idx].deleted {
		return ""
	}

	return hashTable.employees[idx].phone
}

func (hashTable *DoubleHashTable) contains(name string) bool {
	idx, _ := hashTable.find(name)

	if idx < 0 || hashTable.employees[idx] == nil {
		return false
	}

	if hashTable.employees[idx].deleted {
		return false
	}

	return true

}

// Make a display showing whether each slice entry is nil.
func (hashTable *DoubleHashTable) dump_concise() {
	// Loop through the slice.
	for i, employee := range hashTable.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else if employee.deleted {
			fmt.Printf("x")
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

func (hashTable *DoubleHashTable) delete(name string) {
	idx, _ := hashTable.find(name)

	if idx >= 0 {
		hashTable.employees[idx].deleted = true
	}
}

// Return the average probe sequence length for the items in the table.
func (hashTable *DoubleHashTable) ave_probe_sequence_length() float32 {
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

// Show this key's probe sequence.
func (hashTable *DoubleHashTable) probe(name string) int {
	// Hash the key.
	hash1 := hash1(name) % hashTable.capacity
	hash2 := hash2(name) % hashTable.capacity
	fmt.Printf("Probing %s (%d, %d)\n", name, hash1, hash2)

	// Keep track of a deleted spot if we find one.
	deleted_index := -1

	// Probe up to hashTable.capacity times.
	for i := 0; i < hashTable.capacity; i++ {
		index := (hash1 + i + hash2) % hashTable.capacity

		fmt.Printf("    %d: ", index)
		if hashTable.employees[index] == nil {
			fmt.Printf("---\n")
		} else if hashTable.employees[index].deleted {
			fmt.Printf("xxx\n")
		} else {
			fmt.Printf("%s\n", hashTable.employees[index].name)
		}

		// If this spot is empty, the value isn't in the table.
		if hashTable.employees[index] == nil {
			// If we found a deleted spot, return its index.
			if deleted_index >= 0 {
				fmt.Printf("    Returning deleted index %d\n", deleted_index)
				return deleted_index
			}

			// Return this index, which holds nil.
			fmt.Printf("    Returning nil index %d\n", index)
			return index
		}

		// If this spot is deleted, remember where it is.
		if hashTable.employees[index].deleted {
			if deleted_index < 0 {
				deleted_index = index
			}
		} else if hashTable.employees[index].name == name {
			// If this cell holds the key, return its data.
			fmt.Printf("    Returning found index %d\n", index)
			return index
		}

		// Otherwise continue the loop.
	}

	// If we get here, then the key is not
	// in the table and the table is full.

	// If we found a deleted spot, return it.
	if deleted_index >= 0 {
		fmt.Printf("    Returning deleted index %d\n", deleted_index)
		return deleted_index
	}

	// There's nowhere to put a new entry.
	fmt.Printf("    Table is full\n")
	return -1
}

func main() {
	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101", false},
		Employee{"Bob Baker", "202-555-0102", false},
		Employee{"Cindy Cant", "202-555-0103", false},
		Employee{"Dan Deever", "202-555-0104", false},
		Employee{"Edwina Eager", "202-555-0105", false},
		Employee{"Fred Franklin", "202-555-0106", false},
		Employee{"Gina Gable", "202-555-0107", false},
	}

	hash_table := NewDoubleHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	hash_table.probe("Hank Hardy")
	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hash_table.delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	hash_table.dump()

	hash_table.probe("Ann Archer")
	hash_table.probe("Bob Baker")
	hash_table.probe("Cindy Cant")
	hash_table.probe("Dan Deever")
	hash_table.probe("Edwina Eager")
	hash_table.probe("Fred Franklin")
	hash_table.probe("Gina Gable")
	hash_table.set("Hank Hardy", "202-555-0108")
	hash_table.probe("Hank Hardy")

	// Look at clustering.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	big_capacity := 1009
	big_hash_table := NewDoubleHashTable(big_capacity)
	num_items := int(float32(big_capacity) * 0.9)
	for i := 0; i < num_items; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		big_hash_table.set(str, str)
	}
	big_hash_table.dump_concise()
	fmt.Printf("Average probe sequence length: %f\n",
		big_hash_table.ave_probe_sequence_length())
}

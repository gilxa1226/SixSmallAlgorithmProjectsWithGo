package main

import "fmt"

type Employee struct {
	name  string
	phone string
}

type ChainingHashTable struct {
	num_buckets int
	buckets     [][]*Employee
}

// Initialize a ChainingHashTable and return a pointer to it
func NewChainingHashTable(num_buckets int) *ChainingHashTable {
	tmpbuckets := make([][]*Employee, num_buckets)
	return &ChainingHashTable{num_buckets: num_buckets, buckets: tmpbuckets}
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

func (hashTable *ChainingHashTable) dump() {
	for idx := 0; idx < hashTable.num_buckets; idx++ {
		fmt.Println("Bucket ", idx)
		if hashTable.buckets[idx] != nil {
			for x, _ := range hashTable.buckets[idx] {
				fmt.Printf("     %s: %s\n", hashTable.buckets[idx][x].name, hashTable.buckets[idx][x].phone)
			}
		}
	}
}

func (hashTable *ChainingHashTable) find(name string) (int, int) {
	hashVal := hash(name) % hashTable.num_buckets
	if hashTable.buckets[hashVal] == nil {
		return hashVal, -1
	}

	for idx := 0; idx < len(hashTable.buckets[hashVal]); idx++ {
		if hashTable.buckets[hashVal][idx].name == name {
			return hashVal, idx
		}
	}

	return hashVal, -1
}

func (hashTable *ChainingHashTable) set(name string, phone string) {
	bucket, idx := hashTable.find(name)

	if idx > -1 {
		hashTable.buckets[bucket][idx].phone = phone
	} else {
		tmp := Employee{name: name, phone: phone}
		hashTable.buckets[bucket] = append(hashTable.buckets[bucket], &tmp)
	}
}

func (hashTable *ChainingHashTable) get(name string) string {
	bucket, idx := hashTable.find(name)
	if idx > -1 {
		return hashTable.buckets[bucket][idx].phone
	}

	return ""
}

func (hashTable *ChainingHashTable) contains(name string) bool {
	_, idx := hashTable.find(name)

	if idx > -1 {
		return true
	} else {
		return false
	}
}

func (hashTable *ChainingHashTable) delete(name string) {
	bucket, idx := hashTable.find(name)

	if idx < 0 {
		return
	} else {
		hashTable.buckets[bucket] = append(hashTable.buckets[bucket][:idx], hashTable.buckets[bucket][idx+1:]...)
	}

	return
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
		Employee{"Herb Henshaw", "202-555-0108"},
		Employee{"Ida Iverson", "202-555-0109"},
		Employee{"Jeb Jacobs", "202-555-0110"},
	}

	hash_table := NewChainingHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hash_table.delete("Dan Deever")
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
}

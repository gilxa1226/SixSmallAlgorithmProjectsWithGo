/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import (
    "fmt"
    "math/rand"
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

func check_sorted(arr []int) {
    length := len(arr)
    
    for idx := 1; idx < length; idx++ {
        if arr[idx-1] > arr[idx] {
            fmt.Println("The array is NOT sorted!")
            return
        }
    }

    fmt.Println("The array is sorted")
}

func partition(arr []int) int { 
    lo := 0
    hi := len(arr) - 1
    j := lo
    
    pivot := arr[hi]
    
    i := lo - 1
    
    for j = lo; j <= hi - 1; j++ {
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
    if (len(arr) < 2) {
        return
    }
    
    p := partition(arr)
    
    quicksort(arr[:p])
    quicksort(arr[p:])
}

func main() {

    rand.Seed(time.Now().UnixNano())

    // Get the number of items and maximum item value.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted array.
    arr := make_random_array(num_items, max)
    print_array(arr, 40)
    fmt.Println()

    // Sort and display the result.
    quicksort(arr)
    print_array(arr, 40)

    // Verify that it's sorted.
    check_sorted(arr)
}
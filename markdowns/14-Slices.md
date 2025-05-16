## 14- Slices

```go
package main

import (
	"fmt"
	"slices"
)

func main() {

	// Create a slice of integers
	var numbers []int
	var numbers1 = []int{1, 2, 3}
	numbers2 := []int{9, 8, 7}
	fmt.Println("Numbers: ", numbers)   // prints []
	fmt.Println("Numbers1: ", numbers1) // prints [1 2 3]
	fmt.Println("Numbers2: ", numbers2) // prints [9 8 7]

	// Create a slice using the make function
	slice := make([]int, 5) // Create a slice of length 5
	fmt.Println(slice)
	
	// Make slice out of an array
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4] // Slice from index 1 to 3 (4 is excluded) prints [2 3 4
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)   // Append 6 and 7 to the slice
	fmt.Println("Slice1: ", slice1) // prints [2 3 4 6 7]

	// make a copy of the slice
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)               // Copy the contents of slice1 to sliceCopy
	fmt.Println("sliceCopy: ", sliceCopy) // prints [2 3 4 6 7]

	// Create a nil slice
	var nilSlice []int // nil slice
	fmt.Println("Nil Slice: ", nilSlice) // prints []

	for i, v := range slice1 {
		fmt.Println("Index: ", i, "Value: ", v) // prints nothing
	}

	fmt.Println("Element at index 3 of slice1: ", slice1[3]) // prints 6

	slice1[3] = 50
	fmt.Println("Element at index 3 of slice1: ", slice1[3]) // prints 50

	// Check if a slice is equal to another slice
	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal")
	} else {
		fmt.Println("slice1 and sliceCopy are not equal")
	}

	// Create multidimensional slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D Slice: ", twoD) // prints [[0 1 2] [1 2 3] [2 3 4]]

	// index slicing
	slice2 := slice1[2:4]           // Slice from index 2 to 3 (4 is excluded)
	fmt.Println("Slice2: ", slice2) // prints [4 6]

	// check capacity of slice and length of slice
	fmt.Println("The capacity of slice1 is: ", cap(slice1)) // prints 5
	fmt.Println("The length of slice1 is: ", len(slice1))   // prints 5

	// create a slice of strings
	strings := []string{"apple", "banana", "cherry"}
	// check if a string is in the slice
	contains := slices.Contains(strings, "banana")
	if contains {
		fmt.Println("The slice contains banana")
	} else {
		fmt.Println("The slice does not contain banana")
	}

}

```
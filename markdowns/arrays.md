## 13-Arrays

```go
package main

import "fmt"

func main() {

	// var arrayName [size]dataType

	var numbers [5]int   // declare an array of integers with size 5
	fmt.Println(numbers) // prints the default values (0 for int)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	// create an array with values
	fruits := [4]string{"apple", "banana", "orange", "grapes"}
	fmt.Println("Fruits array: ", fruits)

	fmt.Println("Third element:", fruits[2])

	// copy of an array
	originalArray1 := [3]int{1, 2, 3}
	copiedArray1 := originalArray1

	copiedArray1[0] = 100 // This will not affect originalArray

	fmt.Println("Original array:", originalArray1)
	fmt.Println("Copied array:", copiedArray1)

	// iterating over an array
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Value:", numbers[i])
	}

	// using range to iterate
	for i, v := range fruits {
		fmt.Println("Index:", i, "Value:", v)
	}

	// discarding the index using the underscore _
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	// length of an array
	fmt.Println("Length of numbers array:", len(numbers))

	// comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Are array1 and array2 equal?", array1 == array2) // true

	// multi-dimensional arrays
	// 2D array
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)

	// using pointer to modify an array
	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int // declare a pointer to an array

	copiedArray = &originalArray // copiedArray points to originalArray
	copiedArray[0] = 100         // This will not affect originalArray

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", *copiedArray) // dereference the pointer to get the value

}

```
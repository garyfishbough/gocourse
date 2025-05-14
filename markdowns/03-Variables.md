## 03-Variables

```go
package main

import "fmt"

// GLOBAL variable
var middleName = "Cane"

func main() {
	var age int
	var name string = "John"
	var name1 = "Jane"

	count := 10
	lastName := "Smith"

	middleName := "Mayor"
	fmt.Println(middleName)

	// Default values:
	// Numeric Types: 0
	// Boolean Types: false
	// String Type: ""
	// Pointers, slices, maps, functions, and structs: nil

	// ---- SCOPE
	//fmt.Println(firstName)

}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}

```
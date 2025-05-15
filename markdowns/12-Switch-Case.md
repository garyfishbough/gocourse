## 12-Switch Case
```go
package main

import (
	"fmt"
)

func main() {
	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple!")
	case "banana":
		fmt.Println("It's a banana!")
	default:
		fmt.Println("It's some other fruit!")
	}

	// Example with multiple cases
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Not a valid day!")
	}

	// Example with a condition
	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or greater")
	}

	// Example with fallthrough

	num := 2
	switch {
	case num > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is equal to 2")
	default:
		fmt.Println("Number is not 2.")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}

// checkType function demonstrates type switch
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case float64:
		fmt.Println("x is a float64")
	case string:
		fmt.Println("x is a string")
	case bool:
		fmt.Println("x is a bool")
	default:
		fmt.Println("x is of unknown type")
	}
}

```
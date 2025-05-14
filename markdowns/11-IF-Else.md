## 10-If Else

```go
package main

import "fmt"

func main() {

	// Simple if statement
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	// if-else statement
	temperature := 25
	if temperature > 30 {
		fmt.Println("It's hot outside.")
	} else {
		fmt.Println("It's cool outside.")
	}

	// if-else if-else statement
	score := 85
	if score >= 90 {
		fmt.Println("You got an A.")
	} else if score >= 80 {
		fmt.Println("You got a B.")
	} else if score >= 70 {
		fmt.Println("You got a C.")
	} else if score >= 60 {
		fmt.Println("You got a D.")
	} else {
		fmt.Println("You got an F.")
	}

	// Nested if-else
	num := 18
	if num%2 == 0 {
		if num%3 == 0 {
			fmt.Println("The number is divisible by both 2 and 3.")
		} else {
			fmt.Println("The number is divisible by 2 but NOT 3.")
		}
	} else {
		fmt.Println("Number is not divisible by 2.")
	}

	// using or and and operators
	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("Either 10 or 5 is even.")
	}

	if 10%2 == 0 && 6%2 == 0 {
		fmt.Println("Both 10 or 6 are even.")
	}
}


```
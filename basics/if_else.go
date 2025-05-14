package main

import "fmt"

//goland:noinspection ALL
func main() {

	age := 25
	if age >= 18 {
		println("You are an adult.")
	}

	temperature := 25
	if temperature > 30 {
		println("It's hot outside.")
	} else {
		println("It's cool outside.")
	}

	score := 85
	if //goland:noinspection ALL
	score >= 90 {
		println("You got an A.")
	} else if score >= 80 {
		println("You got a B.")
	} else if score >= 70 {
		println("You got a C.")
	} else if score >= 60 {
		println("You got a D.")
	} else {
		println("You got an F.")
	}

	// Nested if-else
	num := 18
	if num%2 == 0 {
		if num%3 == 0 {
			println("The number is divisible by both 2 and 3.")
		} else {
			println("The number is divisible by 2 but NOT 3.")
		}
	} else {
		fmt.Println("Number is not divisible by 2.")
	}

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("Either 10 or 5 is even.")
	}

	if 10%2 == 0 && 6%2 == 0 {
		fmt.Println("Both 10 or 6 are even.")
	}
}

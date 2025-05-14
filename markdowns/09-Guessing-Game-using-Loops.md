## 10-Guessing Game using Loops

```go
package main

import "fmt"
import "math/rand"
import "time"

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1

	var guess int
	attempts := 0

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100. Can you guess it?")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You've guessed the number %d in %d attempts.\n", target, attempts)
			break
		}
	}
}

```
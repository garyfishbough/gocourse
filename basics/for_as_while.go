package main

import "fmt"

func main() {

	// for has while with break
	sum := 0

	for {
		sum += 10
		fmt.Println("Sum: ", sum)
		if sum >= 50 {
			break
		}
	}

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd number: ", num)
		num++
	}

}

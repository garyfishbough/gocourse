package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case
	// Eg. user_id, first_name, http_request

	// UPPERCASE
	// Use case is Constants

	// mixedCase
	// Eg. javaScript, htmlDocument, isValid
	// Starts with lower case

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeeID: ", employeeID)

}

package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//PascalCase
	// Eg CalculateArea, UserInfo, NewHTTPRequest
	// Structs, Interface, Enums

	// snake_case (if used for files)
	// Eg :- user_id, first_name, http_request

	// UPPERCASE
	// EG:- CONSTANTS IN ALL CAPS
	// CONSTANTS :- IMMUTABLE

	//mixedCase
	// eg :- javaScript, htmlDocument, isValid

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("Employee ID: ", employeeID)
}

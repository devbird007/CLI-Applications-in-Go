package main

import "fmt"

// Define the key struct
type TestCaseKey struct {
	Name string
	ID   int
}

func main() {
	// Declare the map
	testCases := map[TestCaseKey]string{}

	// Initialize the map with some test cases
	testCases[TestCaseKey{"AddNumbers", 1}] = "2 + 2 = 4"
	testCases[TestCaseKey{"GreetUser", 2}] = "Hello, World!"
	testCases[TestCaseKey{"ValidateEmail", 3}] = "john.doe@example.com"

	// Access and print a value from the map
	value := testCases[TestCaseKey{"GreetUser", 1}]
	fmt.Println(value) // Output: Hello, World!
}

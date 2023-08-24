package main

import "fmt"

// outside method not allowed - walrus operator
// jwtToken := 30000

const LoginToken string = "abcdef" // First letter capital indicates public variable

func main() {
	var username string = "pulkit"

	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.554455445
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int // default value is 0
	fmt.Println(anotherVariable)

	var anotherString string
	fmt.Println(anotherString) // empty string

	// implicit type
	var website = "pulkit.batra.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}

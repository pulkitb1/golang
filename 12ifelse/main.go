package main

import "fmt"

func main() {
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount >= 10 && loginCount < 20 {
		result = "Watch out"
	} else {
		result = "Danger"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}
}

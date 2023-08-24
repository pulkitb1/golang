package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the watch:")

	// comma ok || err ok
	// no try catch in GO, everything in variable true/false

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating: ", input)
	fmt.Printf("Type of this rating is %T: ", input)
}

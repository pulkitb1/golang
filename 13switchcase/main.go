package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
		fallthrough
	default:
		fmt.Println("case default")
	}
}

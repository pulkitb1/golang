package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	// "time"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in go lang")
	// var numberOne int = 2
	// var numberTwo float32 = 4.5

	// fmt.Println("The sum is:", numberOne+int(numberTwo))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}

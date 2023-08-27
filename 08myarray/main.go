package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[2] = "pineapple"
	fruitList[3] = "grape"

	fmt.Println("Array: ", fruitList)

	var nums [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array: ", nums)

}

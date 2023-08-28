package main

import "fmt"

func main() {
	greeter()
	result := adder(3, 4)
	fmt.Println(result)
	fmt.Println(proAdder(1, 2, 3, 4, 5))
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}
func greeter() {
	fmt.Println("Hi")
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

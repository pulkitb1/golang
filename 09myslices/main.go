package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"apple", "grape", "mango"}
	fmt.Printf("Type of fruitlist is: %T\n", fruitList)

	fruitList = append(fruitList, "orange")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 2
	highScores[1] = 1

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from a slice based on index

	var courses = []string{"reactjs", "javascript", "swift"}
	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

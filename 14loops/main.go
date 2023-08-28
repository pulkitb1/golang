package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday"}
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for day := range days {
	// 	fmt.Println(days[day])
	// }

	for _, day := range days {
		fmt.Println(day)
	}

	i := 0
	for i < 2 {
		fmt.Println(i)
		i++
	}
}

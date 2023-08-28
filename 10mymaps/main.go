package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["ts"] = "typescript"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	delete(languages, "ts")
	fmt.Println(languages)

	value, isPresent := languages["JS"]
	fmt.Println(value, isPresent)

	// loops in maps
	for key, value := range languages {
		fmt.Printf("Key: %v , Value: %v\n", key, value)
	}

}

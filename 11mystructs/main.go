package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	// no inheritance
	pulkit := User{"Pulkit", "pulkit.b@x.com", true, 16}
	fmt.Println(pulkit)
	fmt.Printf("Struct: %+v", pulkit)
}

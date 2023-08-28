package main

import "fmt"

func main() {
	// no inheritance
	pulkit := User{"Pulkit", "pulkit.b@x.com", true, 16, 1}
	fmt.Println(pulkit)
	fmt.Printf("Struct: %+v\n", pulkit)
	pulkit.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println(u.Status)
}

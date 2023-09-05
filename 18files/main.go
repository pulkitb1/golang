package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "this needs to go in a file"
	file, err := os.Create("./myfile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./myfile.txt")

}

func readFile(fileName string) {
	databytes, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Raw data: ", databytes)
	fmt.Println("Data: ", string(databytes))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

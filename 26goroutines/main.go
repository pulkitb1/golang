package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer
var mutex sync.Mutex  // pointer

func main() {
	// go greeter("hello")
	// greeter("world")
	websiteList := []string{
		"https://loc.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Endpoint invalid")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

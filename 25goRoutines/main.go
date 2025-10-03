package main

import (
	"fmt"
	"net/http"
	"sync"
	//"time"
)

var wg sync.WaitGroup

func main() {
	//go greeter("hello") // Goroutine is introduced through the keyword go->This creates another thread
	//greeter("world")
	defer wg.Done()
	websitelist := []string{
		"https://dummyjson.com/products",
		"https://dummyjson.com/products/1",
		"https://dummyjson.com/products/3",
		"https://dummyjson.com/products/4",
		"https://dummyjson.com/products/5",
		"https://fb.com/products",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}
/*
func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
*/
func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Printf("%d is status code for %s\n", res.StatusCode, endpoint)
}

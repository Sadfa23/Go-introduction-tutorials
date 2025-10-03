package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummyjson.com/products/1"

func main() {
	fmt.Println("Web Requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T", response)

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err!=nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println("This is the content", content)

	defer response.Body.Close() // caller's responsibility to close connection
}

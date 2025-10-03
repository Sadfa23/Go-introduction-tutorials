package main

import (
	"fmt"
	"net/url"
)
const myUrl string = "https://dummyjson.com/products/search?q=phone&paymentid=746r"
func main() {
	fmt.Println("Lesson on url handling")
	fmt.Println(myUrl)

	//parsing url
	result, _:= url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params are %T", qparams)

	fmt.Println(qparams["q"])
	for _, val := range qparams {
		fmt.Println("Param is ", val)
	}

	// creatting a url from its components
	partsOfUrl := &url.URL { // The & ensures a ref of the url not a copy
		Scheme: "https",
		Host: "lco.dev",
		Path: "tutor",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}

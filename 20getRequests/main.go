package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Get Requests with Go")

	//PerformGetRequest()
	//PerformPostJSONRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://dummyjson.com/products/1"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code ", response.StatusCode)
	//fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String()) //Alternative
	// fmt.Println(string(content))
}

func PerformPostJSONRequest () {
	myUrl := "https://dummyjson.com/products/add"

	requestBody := strings.NewReader(`
	{
	"title": "Jordan"
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err!= nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformFormRequest () {
	myUrl := "https://dummyjson.com/products/add"
	// Form data

	data := url.Values{}
	data.Add("FirstName", "Terence")
	data.Add("MiddleName", "Sadfa")
	data.Add("SurName", "Otieno")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

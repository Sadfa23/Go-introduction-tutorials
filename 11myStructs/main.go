package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// There are no super, inheritance, parent etc in Go
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 60}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	fmt.Printf("Hitesh name is %v and email is %v", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

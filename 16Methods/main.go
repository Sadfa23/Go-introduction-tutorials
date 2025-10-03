package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// There are no super, inheritance, parent etc in Go
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 60}
	//fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	//fmt.Printf("Hitesh name is %v and email is %v", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status, "\n")
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is", u.Email)
}

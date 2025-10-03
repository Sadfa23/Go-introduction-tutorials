package main 
import "fmt"

func main () {
	fmt.Println("Welcome to the class for pointers")

	//var ptr *int // * denotes pointer
	//fmt.Println("Value of ponter is", ptr)

	myNumber := 32

	var ptr = &myNumber  // & references

	fmt.Println("Value of pointer is", ptr) // gives the memory address
	fmt.Println("Value of actual pointer is", *ptr) // gives the value 32

	*ptr = *ptr *2
	fmt.Println("New value is: ", myNumber)
}
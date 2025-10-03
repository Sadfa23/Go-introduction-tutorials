package main

import "fmt"

func main() {
	fmt.Println("Lesson on functions in golang")
	greeter()	
	result:= adder(3, 5)
	fmt.Println("Result is:", result)

	proRes, myMessage := proAdder(2,3,4,5,6,7,8,9)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro result message: ", myMessage)
}

func greeter() {
	fmt.Println("Hello Terence from Golang")
}

func adder(valOne int, valTwo int) int {
	return valOne * valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return  total, "This is the final result"
}
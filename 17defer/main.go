package main

import "fmt"

func main() {
	
	defer fmt.Println("Defered program-1") // this is defered to a later time
	defer fmt.Println("Defered program-2") // this is defered to a later time
	defer fmt.Println("Defered program-3") 
	defer fmt.Println("Defered program-4") 
	// 4, 3, 2, 1

	fmt.Println("Lesson on defer") // Last in, First Out
	
	myDefer()
}

func myDefer() {
	 for i:=0; i< 5 ; i++ {
		defer fmt.Println(i)
	 }
}
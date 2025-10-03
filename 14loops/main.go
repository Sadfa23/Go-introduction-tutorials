package main

import "fmt"

func main() {
	fmt.Println("Lesson on loops")

	//days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
	//fmt.Println(days)
	/*

	for d:=0; d <len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %x\n", index, day)
	}  
	*/

	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue ==7 {
			goto lco
		}
		
		if rogueValue == 5 {
			break
		}

		
		/*
		if rogueValue == 5 {
			rogueValue ++ //Outputs 1 2 3 4 6 7 8 9
			continue
		*/
		fmt.Println("Value is :", rogueValue)
		rogueValue++
	}

	lco: 
		fmt.Println("Jumping into .com")

}
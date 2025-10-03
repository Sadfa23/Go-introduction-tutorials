package main

import (
	"fmt"
	"sort"
)
func main () {
	fmt.Println("Welcome to lesson on slices")
	var fruitList = []string{"Apple", "Tomateo", "Peach"}

	fmt.Printf("Type of fruiList is %T", fruitList)

	fruitList =  append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 255
	highScore[1] = 254
	highScore[2] = 297
	highScore[3] = 289

	highScore = append(highScore, 555, 666, 321) // memory is re-allocated from 4 to the new length

	fmt.Println(highScore)
	fmt.Print(sort.IntsAreSorted(highScore))

	sort.Ints(highScore)
	fmt.Println(highScore)

	//How to remove a value from slices based on index

	var courses = []string{"reactjs", "JS","TS", "Python", "ruby"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1 :]... )
	fmt.Println(courses)

}


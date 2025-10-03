package main

import "fmt" 

func main () {
	fmt.Println("Welcome to array in golang")

	var fruits [4]string
	fruits[0] = "Tomato"
	fruits[1] = "Peach"
	fruits[3] = "Pear"

	fmt.Println("Fruit list", fruits)
	fmt.Println("Fruit list",len(fruits))

	var vegList = [4]string{"potata", "mango", "peach"}

	fmt.Println("Veg list",len(vegList))


}
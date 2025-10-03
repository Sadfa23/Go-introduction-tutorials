package main 

import "fmt"
func main () {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["Go"] = "Golang"
	languages["TS"] = "Typescript"
	languages["rb"] = "ruby"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS is short for", languages["JS"])

	delete(languages, "rb")
	fmt.Println("List of all languages", languages)

	//Loops intro
	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value )
	}
}
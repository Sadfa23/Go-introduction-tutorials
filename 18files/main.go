package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Go")

	content := "This is content to a file in this ditectory"

	file, err := os.Create("./localFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is", length)
	readFile("./localFile.txt")
	defer file.Close()
}


func readFile(filename string) {
	dataInByteForm, err := ioutil.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("Text data inside the file is", dataInByteForm)
	fmt.Println("Text data inside the file is", string(dataInByteForm))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
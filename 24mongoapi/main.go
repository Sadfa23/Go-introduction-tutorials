package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
	
)

func main() {
	fmt.Println("Server is getting started")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4005", r))
	fmt.Println("Listening to port 4005")
}

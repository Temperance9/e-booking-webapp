package main

import (
	"e-book_webapp/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes := routes.SetupRouter()
	fmt.Println("we are live ")
	log.Fatal(http.ListenAndServe(":8080", routes))
}

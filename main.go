package main

import (
	user "e-book_webapp/services"
	"e-book_webapp/utils"
	"fmt"
)

func main() {
	/*
		routes := routes.SetupRouter()
		fmt.Println("we are live ")
		log.Fatal(http.ListenAndServe(":8080", routes))
	*/
	db, err := utils.ConnectDatabase()
	defer db.Close()
	if err != nil {
		fmt.Println("el db maconnectetech !!")
	}
	/*
		u := user.User{Username: "fizou", Password: "12313", Email: "aymenfaz@gmail.com"}
		i, err := user.CreateUser(db, u)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(i)
	*/

	u, err := user.GetUser(db, 1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(u)
}

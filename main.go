package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"log"
	"github.com/bernhardkern/helloworld_go/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/bernhardkern/helloworld_go/rest"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello world3")

	mySQLRepo, err := data.NewMySQL()
	if err != nil {
		fmt.Println("bad error", err)
	}
	restHandler := rest.NewRestHandler(mySQLRepo)
	router := httprouter.New()

	router.GET("/person", restHandler.HandleGet)
	router.POST("/person", restHandler.HandlePost)

	person := data.Person{Id: 1, FirstName: "B", LastName: "K"}
	person2 := data.Person{Id: 2, FirstName: "Johannes", LastName: "Brüderl"}
	err = mySQLRepo.Create(person)
	if err != nil {
		fmt.Println("bad error", err)
	}
	err = mySQLRepo.Create(person2)
	if err != nil {
		fmt.Println("bad error", err)
	}
	fmt.Println("finished successfully")

	log.Fatal(http.ListenAndServe(":8888", router))
}

func square(number int) int {
	return number * number
}

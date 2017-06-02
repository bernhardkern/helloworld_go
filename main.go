package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/bernhardkern/helloworld_go/data"
	"github.com/bernhardkern/helloworld_go/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mySQLRepo, err := data.NewMySQL("test", "test", "test")
	if err != nil {
		fmt.Println("bad error", err)
	}
	restHandler := rest.NewRestHandler(mySQLRepo)
	router := httprouter.New()

	router.GET("/person", restHandler.HandleGet)
	router.POST("/person", restHandler.HandlePost)

	log.Print("Starting server")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func square(number int) int {
	return number * number
}

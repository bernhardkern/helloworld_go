package main

import (
	"fmt"
	"net/http"

	"log"

	"flag"

	"github.com/bernhardkern/helloworld_go/data"
	"github.com/bernhardkern/helloworld_go/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var (
	mySQLUser     string
	mySQLPassword string
	mySQLSchema   string
)

func init() {
	flag.StringVar(&mySQLUser, "user", "test", "MySQL user name")
	flag.StringVar(&mySQLPassword, "password", "test", "MySQL password")
	flag.StringVar(&mySQLSchema, "schema", "test", "MySQL schema")
}

func main() {
	flag.Parse()

	mySQLRepo, err := data.NewMySQL(mySQLUser, mySQLPassword, mySQLSchema)
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

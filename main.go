package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"log"
	"github.com/bernhardkern/helloworld_go/data"
	_ "github.com/go-sql-driver/mysql"
)

func Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello world3")
	router := httprouter.New()
	router.GET("/", Test)

	mySql, err := data.NewMySQL()
	if err != nil {
		fmt.Println("bad error", err)
	}
	person := data.Person{Id: 1, FirstName: "B", LastName: "K"}
	person2 := data.Person{Id: 2, FirstName: "Johannes", LastName: "Brüderl"}
	err = mySql.Create(person)
	if err != nil {
		fmt.Println("bad error", err)
	}
	err = mySql.Create(person2)
	if err != nil {
		fmt.Println("bad error", err)
	}
	fmt.Println("finished successfully")

	log.Fatal(http.ListenAndServe(":8888", router))

}

func square(number int) int {
	return number * number
}

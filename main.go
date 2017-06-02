package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello world3")
	router := httprouter.New()
	router.GET("/", Test)
	log.Fatal(http.ListenAndServe(":8888", router))
}

func square(number int) int {
	return number * number
}

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the RESTful api")
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)

	log.Println("Running api server")

	http.ListenAndServe(":8787", router)
}

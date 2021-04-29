package main

import (
	"fmt"
	"go-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {

	_router := mux.NewRouter().StrictSlash(true)
	
	_router.HandleFunc("/", controllers.HomePage)
	_router.HandleFunc("/article", controllers.CreateArticle).Methods("POST")
	_router.HandleFunc("/article", controllers.GetArticle).Methods("GET")


	fmt.Println("http://localhost:10000/")

	log.Fatal(http.ListenAndServe(":10000", _router))
}

func main() {
	handleRequests()
}
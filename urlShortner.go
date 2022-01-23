package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UrlStruct struct {
//	structure for url shortening logic handling
}

var urlStructs []UrlStruct



func shortUrl(w http.ResponseWriter, r *http.Request) {
//	main function for handling url shortning
}

func returnAllShortenedUrl(w http.ResponseWriter, r *http.Request){
//	function for returning all previous shortened url's

}

func apiRequests() {
	//api request definitions for handling url shortning
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/shortUrl", shortUrl).Methods("POST")
	route.HandleFunc("/getUrl", returnAllShortenedUrl).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", route))
}


func main() {
	//main function
	apiRequests()
}

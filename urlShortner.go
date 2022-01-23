package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type UrlStruct struct {
//	structure for url shortening logic handling
	OriginalUrl string `json:"originalUrl"`
	ShortenedUrl string `json:"shortenedUrl"`
	Id string `json:"id"`
}

var urlStructs []UrlStruct



func shortUrl(w http.ResponseWriter, r *http.Request) {
//	main function for handling url shortning
	reqBody, _ := ioutil.ReadAll(r.Body)
	var urlStruct UrlStruct
	err :=json.Unmarshal(reqBody, &urlStruct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error occur when marshalling request data: %v\n", err)
	}else if urlStruct.OriginalUrl == ""{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error occur when marshalling request data: %v\n", err)
		return
	}else {
		for _, url := range urlStructs {
			if url.OriginalUrl == urlStruct.OriginalUrl {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w,"Url %s is already shortened, and Shortened url is %s\n", url.OriginalUrl, url.ShortenedUrl)
				return
			}
		}
		//use of hashing logic  to  generating id and  appending  it to short url
		hd := hashids.NewData()
		h, _ :=  hashids.NewWithData(hd)
		//using current time in unix timestamp format for creating unique id for short url
		now := time.Now()
		urlStruct.Id, _ = h.Encode([]int{int(now.Unix())})
		urlStruct.ShortenedUrl =  "http://localhost:3000/" + urlStruct.Id
		urlStructs = append(urlStructs, urlStruct)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Printf("Successfully Shortened  the Url:%v",urlStructs)
		json.NewEncoder(w).Encode(urlStruct)
	}
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

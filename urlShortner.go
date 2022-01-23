package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type UrlStruct struct {
//	structure for url shortening logic handling
	OriginalUrl string `json:"originalUrl"`
	ShortenedUrl string `json:"shortenedUrl"`
	Id string `json:"id"`
}

var urlStructs []UrlStruct


var INPUTFILEFORETLDATAPROCESSING = "url.json"

func writeFile(b []byte) bool {
	err3 := ioutil.WriteFile(INPUTFILEFORETLDATAPROCESSING, b, 0644)
	if err3 != nil {
		//fmt.Println("Error occur when writing to a file: %s", err3)
		return false
	}
	return true
}
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
		fmt.Printf("originalUrl data is empty:\n")
		return
	}else {
		for _, url := range urlStructs {
			if url.OriginalUrl == urlStruct.OriginalUrl {
				w.WriteHeader(http.StatusOK)
				fmt.Printf("Url %s is already shortened, and Shortened url is %s\n", url.OriginalUrl, url.ShortenedUrl)
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
		b, err := json.MarshalIndent(urlStructs,"","\t")
		if err != nil {
			fmt.Printf( "Error occur when marshalling request data: %v", err)
		}
		if writeFile(b) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Printf("Successfully Shortened  the Url:%v\n", urlStruct)
			json.NewEncoder(w).Encode(urlStruct)
		}else{
			w.WriteHeader(http.StatusOK)
			fmt.Printf("Successfully Shortened Url,But Got Error  While  Updated Storage File  %s:\n", urlStruct.OriginalUrl)
			json.NewEncoder(w).Encode(urlStruct)
		}
	}
}

func getUrByld(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	for _, url := range urlStructs {
		if url.Id == key {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(url)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func returnAllShortenedUrl(w http.ResponseWriter, r *http.Request){
//	function for returning all previous shortened url's
	json.NewEncoder(w).Encode(urlStructs)
}

func apiRequests() {
	//api request definitions for handling url shortning
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", shortUrl).Methods("POST")
	route.HandleFunc("/", returnAllShortenedUrl).Methods("GET")
	route.HandleFunc("/url/{id}", getUrByld).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", route))
}


func main() {
	//main function
	if _, err := os.Stat(INPUTFILEFORETLDATAPROCESSING); err == nil {
		content, err1 := ioutil.ReadFile(INPUTFILEFORETLDATAPROCESSING)
		if err1 != nil {
			fmt.Println("Error when opening storage file for short and original url's: ", err1)
		}
		if len(content)>0 {
			err2 := json.Unmarshal(content, &urlStructs)
			if err2 != nil {
				fmt.Println("Error while unmarshaling contents of  storage file for short and original url's: ", err2)
			}
		}
	}
	apiRequests()
}






package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html")

	w.Write([]byte("<h1 style='color: steelblue'> Hello </h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := Book{
		Title:  "The Gunslinger",
		Author: "stephen King",
		Pages:  304,
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

package main

import (
	_ "filter"
	"frame/filter"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", filter.Exe)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

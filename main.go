package main

import (
//	"encoding/json"
//	"fmt"
//	"log"
	"net/http"
//	"github.com/PuerkitoBio/goquery"
)

	const portNumber = ":8080"

	func main() {
		http.HandleFunc("/users", UsersHandler)
		http.HandleFunc("/wiki", wikiHandler)
		http.ListenAndServe(portNumber, nil)
	}
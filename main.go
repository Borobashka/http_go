package main

import (
//	"encoding/json"
//	"fmt"
//	"log"
	"net/http"
	"http_go-master/handler"
//	"github.com/PuerkitoBio/goquery"
)

	const portNumber = ":8080"

	func main() {
		http.HandleFunc("/users", UsersHandler)
		http.HandleFunc("/wiki", WikiHandler)
		http.ListenAndServe(portNumber, nil)
	}
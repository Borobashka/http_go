package main

import (
//	"encoding/json"
//	"fmt"
//	"log"
	"net/http"
	"github.com/Borobashka/http_go/handler"
//	"github.com/PuerkitoBio/goquery"
)

	const portNumber = ":8080"

	func main() {
		http.HandleFunc("/users", handler.UsersHandler)
		http.HandleFunc("/wiki", handler.UsersHandler)
		http.ListenAndServe(portNumber, nil)
	}
package main

import (
//	"encoding/json"
//	"fmt"
//	"log"
	"net/http"
	"github.com/Borobashka/http_go/v1"
//	"github.com/PuerkitoBio/goquery"
)

	const portNumber = ":8080"

	func main() {
		http.HandleFunc("/users", UsersHandler)
		http.HandleFunc("/wiki", WikiHandler)
		http.ListenAndServe(portNumber, nil)
	}
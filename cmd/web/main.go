package main

import (
	"fmt"
	"net/http"

	"github.com/Kevinchayadi/go-modern-app/pkg/handlers"
)

var portnumber = ":8080"





func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	
	fmt.Println(fmt.Sprintf("Starting application with port %s", portnumber))

	_ = http.ListenAndServe(portnumber, nil)
}
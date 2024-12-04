package main

import (
	
	"fmt"
	"net/http"
)

var portnumber = ":8080"

func home(w http.ResponseWriter,r *http.Request){
		// fmt.Fprintf(w, "this is the home page")
}

func about(w http.ResponseWriter, r *http.Request) {
	// sum:= AddValue(2,2)
    // _, _= fmt.Fprintf(w, fmt.Sprintf("this is the about page, two plus two is %d",sum))
}



func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	
	fmt.Println(fmt.Sprintf("Starting application with port %s", portnumber))

	_ = http.ListenAndServe(portnumber, nil)
}
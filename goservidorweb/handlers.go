package main

import (
	"fmt"
	"net/http"
)


func HandleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Server is again alive")
}

func HandleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the API endpoint")
}
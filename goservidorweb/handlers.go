package main

import (
	"fmt"
	"net/http"
)


func HandleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Server is again alive")
}
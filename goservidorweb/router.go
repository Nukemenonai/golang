package main

import (
	"fmt"
	"net/http"
)

type Router struct{
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router{
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) Findhandler(path string) (http.HandlerFunc, bool){
	handler, exist := r.rules[path]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "server is alive")
}
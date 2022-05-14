package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	//recibe un handler y lo retorna si su logica pasa, o una vez ejecute su logica
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("checking authentication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}

}
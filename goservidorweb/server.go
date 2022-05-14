package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port string 
	router *Router
}


//se devuelve el puntero al servidor porque no queremos una copia sino un servidor que pueda ser modificado
func NewServer(port string) *Server{
	return &Server{
		port: port,
		router: NewRouter(),
	}

}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		fmt.Println("error!!")
		return err
	}
	return nil
}
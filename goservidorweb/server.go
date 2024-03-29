package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

//se devuelve el puntero al servidor porque no queremos una copia sino un servidor que pueda ser modificado
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}

}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}


func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, m := range middlewares {
		f = m(f)
	}
	return f

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

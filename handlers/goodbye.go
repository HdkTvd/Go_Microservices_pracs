package handlers

import (
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (g *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Good bye message")
	rw.Write([]byte("Good Bye"))
}

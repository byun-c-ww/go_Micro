package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func New_Goodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (b *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b.l.Print("hello from goodbye.go")
	rw.Write([]byte("goodbye user"))
}

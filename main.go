package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "package-log:", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.HandleFunc("/", hh)

	http.ListenAndServe("127.0.0.1:9090", sm)

}

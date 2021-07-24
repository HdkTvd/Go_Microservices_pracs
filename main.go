package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/hello/handlers"
)

func main() {

	l := log.New(os.Stdout, "package-log:", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", gh)
	sm.Handle("/product", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, Graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

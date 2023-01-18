package main

import (
	"fmt"
	"log"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello :)")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func MyMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recieved a %s request for the source:  %s", r.Method, r.URL)
		handler.ServeHTTP(w, r)
		log.Println("Handler finished processing request")
	})
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/status/", StatusHandler)

	WrappedMux := MyMiddleware(mux)

	go http.ListenAndServe(":8080", WrappedMux)
}

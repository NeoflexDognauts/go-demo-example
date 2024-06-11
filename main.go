package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
}

func healzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthly")
}

func main() {
	go serveHealz()
	serveMain()
}

func serveMain() {
	http.HandleFunc("/test", handler)
	http.HandleFunc("/test2", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHealz() {
	http.HandleFunc("/healz", healzHandler)
	log.Fatal(http.ListenAndServe(":8089", nil))
}

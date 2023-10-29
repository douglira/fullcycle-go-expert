package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-ctx.Done():
		msg := "Request cancelada pelo cliente"
		log.Println(msg)
	case <-time.After(5 * time.Second):
		msg := "Request processada com sucesso"
		log.Println(msg)
		w.Write([]byte(msg))
	}
}

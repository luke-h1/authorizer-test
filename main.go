package main

import (
	"errors"
	"log"
	"net/http"
	"os"
)

func OnlyErrors() error {
    return errors.New("something went wrong")
}

func handler(w http.ResponseWriter, r *http.Request) {
    authorizationValue := os.Getenv("API_KEY")

    sourceIp := r.Header.Get("x-forwarded-for")
    consumer := r.Header.Get("x-consumer")

    log.Printf("Auth header: %s", r.Header.Get("authorization"))

    for k, v := range r.Header {
        log.Printf("Header: %s = %s", k, v[0])
    }

    if r.Header.Get("authorization") == authorizationValue {
        log.Printf("Authorized request from %s, consumer: %s", sourceIp, consumer)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Authorized"))
        return
    }
    log.Printf("Unauthorized request from %s, consumer: %s", sourceIp, consumer)
    w.WriteHeader(http.StatusUnauthorized)
    w.Write([]byte("Unauthorized"))
}

func init() {
    if os.Getenv("API_KEY") == "" {
        log.Fatal("API_KEY environment variable is required")
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
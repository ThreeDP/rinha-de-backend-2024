package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/ThreeDP/rinha-de-backend/route"
)

const PORT string = ":8080"

func main() {
	db := route.DBQueries{}
	server := &route.BankServer{Store: &db}
	if err := http.ListenAndServe(PORT, server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
	fmt.Printf("Listen on Port: %s\n", PORT)
}
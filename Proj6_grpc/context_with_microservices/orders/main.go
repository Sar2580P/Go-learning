package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main(){
	router:= http.NewServeMux()
	router.HandleFunc("GET /users/{id}/products", handleGetUserProducts)

	server:= http.Server{
		Addr: ":8081",
		Handler: router,
	}

	log.Println("Starting server on port 8081")
	log.Fatal(server.ListenAndServe())
}

func handleGetUserProducts(w http.ResponseWriter, r *http.Request){
	userID:= r.PathValue("id")   // query-params

	log.Printf("Getting products for user %s...", userID)

	// get user products from DB
	time.Sleep(4*time.Second)
	products:= []string{"product1", "product2"}
	writeJSON(w, http.StatusOK, products)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)


type User struct{
	ID string 
	Name string 
	Products []string
}

func main(){

	// creating simple rest api 
	router := http.NewServeMux()

	// this api interacts with 2 microservices: get-user , get-products
	router.HandleFunc("GET /users/{id}", handleGetUserProducts)

	server:= http.Server{
		Addr: ":8080", 
		Handler: router, 
	}

	log.Println("Starting server on port 8080")
	log.Fatal(server.ListenAndServe())
}

func handleGetUserProducts(w http.ResponseWriter, r *http.Request){
	userID:= r.PathValue("id")  // query-param

	// get user info from DB...
	user:= User{
		ID: userID,
		Name: "John Doe", 
		Products: []string{},
	}

	// ☀️ context 
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel() 


	url:= fmt.Sprintf("http://localhost:8001/users/%s/products", userID)
	req, err:= http.NewRequestWithContext(ctx, "GET", url, nil)
	if err!=nil{
		// writing response
		http.Error(w, "foo", http.StatusInternalServerError)
		return
	}

	resp, err:= http.DefaultClient.Do(req)   // interacting with get-user microservice
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	defer resp.Body.Close()

	// data is array of strings
	data, err:= io.ReadAll(resp.Body)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// unmarshal products as bytes into user.Products
	if err:= json.Unmarshal(data, &user.Products); err!=nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return 
	}
	writeJSON(w, http.StatusOK, user)

}


func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}


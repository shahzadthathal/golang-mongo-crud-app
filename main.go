package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Api Base Path

	//Routes

	//HandleFunc registers a new route with a matcher for the URL path.
	s.HandleFunc("/createProfile", createProfile).Methods("Post")
	s.HandleFunc("/get/all/users", getAllUsers).Methods("GET")
	s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")
	s.HandleFunc("/delete/profile/{id}", deleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", s)) //Run Server
}

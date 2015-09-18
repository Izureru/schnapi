package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Index was called")
	fmt.Fprintln(w, "Welcome!")
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Name: "john doe"},
		User{Name: "jane doe"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	log.Println("User Show was called")
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Fprintln(w, "User show:", userId)
}

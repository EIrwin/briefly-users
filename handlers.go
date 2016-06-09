package main

import (
	"encoding/json"
	"net/http"
    "io"
    "io/ioutil"
	"github.com/gorilla/mux"
	"log"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Pong"); err != nil {
		panic(err)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {	

	vars := mux.Vars(r)
	id := vars["id"]
	
	user,err := GetUser(id)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var user User
	
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	u,err := CreateUser(user)
	if err != nil {
		msg := err.Error()
		log.Fatal(msg)
		w.WriteHeader(http.StatusBadRequest)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

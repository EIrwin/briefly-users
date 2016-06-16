package api

import (
	"encoding/json"
	"net/http"
    	"io"
    	"io/ioutil"
	"github.com/gorilla/mux"
	"log"

	"github.com/eirwin/briefly-users/services"
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

	log.Print(id)
	req := services.GetUserRequest{Id:id}
	user,err := services.GetUser(&req)
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
	req := services.CreateUserRequest{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	
	if err := json.Unmarshal(body, &req); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	user,err := services.CreateUser(&req)
	if err != nil {
		msg := err.Error()
		log.Fatal(msg)
		w.WriteHeader(http.StatusBadRequest)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

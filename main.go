package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name string `json:"name"`
}

var Todos []*Todo

const baseURL = "0.0.0.0:8080"

func main() {
	// http.HandleFunc("/hello", Get)

	r := mux.NewRouter()
	//r.HandleFunc("/todos", Create).Methods(http.MethodPost)
	r.HandleFunc("/todos", Create).Methods(http.MethodPost)
	r.HandleFunc("/todos", Get).Methods(http.MethodGet)

	log.Println("Listening in url " + baseURL)
	log.Fatal(http.ListenAndServe(baseURL, r))
}

func PrintHello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("hello world"))
	}

	if _, err := w.Write([]byte("hello world")); err != nil {
		log.Println(err)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	Todos = []*Todo{{
		Name: "1",
	}, {
		Name: "99",
	}}
	todosRes, _ := json.Marshal(Todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(todosRes)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)
	Todos = append(Todos, &t)
	w.Write([]byte("Success add todo " + t.Name))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func GameIndex(w http.ResponseWriter, r *http.Request) {
	games := Games{
		Game{Name: "Tic Tac Toe"},
		Game{Name: "Checkers"},
	}
	if err := json.NewEncoder(w).Encode(games); err != nil {
		panic(err)
	}
}

func GameShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	fmt.Fprintf(w, "Game show: %s\n", gameID)
}

func GetMove(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// gameID := vars["gameID"]
	fmt.Println("we here")
	json := simplejson.New()
	pos := []int{0, 0}
	json.Set("move", pos)
	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

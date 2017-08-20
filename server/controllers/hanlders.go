package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	gameId := vars["gameId"]
	fmt.Fprintf(w, "Game show: %s\n", gameId)
}

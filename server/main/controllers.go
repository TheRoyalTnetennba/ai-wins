package main

import (
	"fmt"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func GameShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	fmt.Fprintf(w, "Game show: %s\n", gameID)
}

func GetMove(w http.ResponseWriter, r *http.Request) {
	req, _ := simplejson.NewFromReader(r.Body)
	ch := make(chan []byte)
	go GetAIMove(req.MustMap(), ch)
	payload := <-ch
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func ListGames(w http.ResponseWriter, r *http.Request) {
	var games []*Game
	keys, err := client.GetAll(ctx, datastore.NewQuery("Game"), &games)
	// for i, key := range keys {
	//     fmt.Println(key)
	//     fmt.Println(games[i])
	// }
	fmt.Println(games)
	w.Header().Set("Content-Type", "application/json")
	
}

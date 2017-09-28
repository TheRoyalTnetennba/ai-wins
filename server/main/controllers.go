package main

import (
	"fmt"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"cloud.google.com/go/datastore"
	"github.com/TheRoyalTnetennba/ai-wins/server/games/ticTacToe"
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

func GamesIndex(w http.ResponseWriter, r *http.Request) {
	res := simplejson.New()
	var games []*Game
	keys, _ := Client.GetAll(Ctx, datastore.NewQuery("Game"), &games)
	for i, key := range keys {
	    fmt.Println(key)
	    fmt.Println(games[i])
	}
	fmt.Println(games)
	res.Set("games", games)
	w.Header().Set("Content-Type", "application/json")
	payload, _ := res.Encode()
	w.Write(payload)
}

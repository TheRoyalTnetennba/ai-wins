package main

import (
	"fmt"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	// "github.com/gorilla/sessions"

	"cloud.google.com/go/datastore"
	"github.com/TheRoyalTnetennba/ai-wins/server/games/ticTacToe"
	"github.com/TheRoyalTnetennba/ai-wins/server/games/wordsWithUnfeelingMachines"
)

func GameShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	fmt.Fprintf(w, "Game show: %s\n", gameID)
}

func GetMove(w http.ResponseWriter, r *http.Request) {
	req, _ := simplejson.NewFromReader(r.Body)
	ch := make(chan []byte)
	go ttt.GetAIMove(req.MustMap(), ch)
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

func WWUFGetLetters(w http.ResponseWriter, r *http.Request) {
    letters := wwuf.NewLetterSet()
    res := simplejson.New()
    // ch := make(chan []byte)
    res.Set("letters", letters)
    payload, _ := res.Encode()
    w.Header().Set("Content-Type", "application/json")
    w.Write(payload)
    w.Write(payload)
}

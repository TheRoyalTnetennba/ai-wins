package main

import (
	"fmt"
	"log"
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
	req, err1 := simplejson.NewFromReader(r.Body)
	res := simplejson.New()
	// gameID := (req.Get("gameID")).MustString()
	// marker := (req.Get("marker")).MustArray()
	gameState := (req.Get("gameState")).MustArray()
	pos := []int{0, 0}
	res.Set("move", pos)
	payload, err := res.Encode()
	if err != nil || err1 != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

package main

import (
    "fmt"
	"net/http"
    "github.com/TheRoyalTnetennba/ai-wins/server/packages/db"
    "github.com/TheRoyalTnetennba/ai-wins/server/packages/ttt"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GoogleLogin",
		"GET",
		"/api/v1/googlelogin",
		GoogleLogin,
	},
	Route{
		"GoogleCallback",
		"GET",
		"/api/v1/googlecallback",
		GoogleCallback,
	},
	Route{
		"GameIndex",
		"GET",
		"/api/v1/games",
		GameIndex,
	},
	Route{
		"UserShow",
		"GET",
		"/api/v1/sec/user",
		UserShow,
	},
	Route{
		"tttState",
		"POST",
		"/api/v1/sec/tic-tac-toe",
		tttState,
	},
	Route{
		"hangmanState",
		"POST",
		"/api/v1/sec/hangman",
		hangmanState,
	},
	Route{
		"wwufState",
		"POST",
		"/api/v1/sec/words-with-unfeeling-machines",
		wwufState,
	},
}

func GameIndex(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go getGames(w, r, c)
    respond(w, c)
}

func UserShow(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    if db.VerifySessionToken(w, r, c) {
    	go getUser(w, r, c)
    	respond(w, c)
    }
}

func tttState(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    fmt.Println("the fuck all that has happened thus far is not the problem")
    if db.VerifySessionToken(w, r, c) {
        fmt.Println("verify session is not the problem")
        ttt.Move(w, r, c)
        respond(w, c)
    } else {
        c <- []byte("no bueno")
        respond(w, c)
    }
}

func hangmanState(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    if verifySessionToken(w, r, c) {
        // go ttt.Move(w, r, c)
        respond(w, c)
    }
}

func wwufState(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    if verifySessionToken(w, r, c) {
        // go ttt.Move(w, r, c)
        respond(w, c)
    }
}
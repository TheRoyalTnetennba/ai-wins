package main

import (
	"net/http"
    "encoding/json"
    "github.com/gorilla/mux"
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
        "GameIndex",
        "GET",
        "/games",
        GameIndex,
    },
    Route{
        "Token",
        "POST",
        "/token",
        Token,
    },
    Route{
        "Login",
        "POST",
        "/login",
        Login,
    },
	Route{
		"GameRoute",
		"POST",
		"/game/{resource}",
		GameRoute,
	},
}

func GameIndex(w http.ResponseWriter, r *http.Request) {
    games := getAllGames()
    respond(w, r, 200, games)
}

func Token(w http.ResponseWriter, r *http.Request) {
    t := User{}
    json.NewDecoder(r.Body).Decode(&t)
    user := getUserBySessionToken(t)
    respond(w, r, 200, user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    gUser := GoogleUser{}
    json.NewDecoder(r.Body).Decode(&gUser)
    user := processGoogleLogin(gUser)
    respond(w, r, 200, user)
}

func GameRoute(w http.ResponseWriter, r *http.Request) {
    c, n := User{}, User{}
    json.NewDecoder(r.Body).Decode(&c)

    switch vars := mux.Vars(r); vars["resource"] {
    case "ttt":
        n = processTTT(c)

    }
    respond(w, r, 200, n)
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    payload, _ := json.Marshal(data)
    w.Write(payload)
}


// func Login(w http.ResponseWriter, r *http.Request) {

// }

// func GameIndex(w http.ResponseWriter, r *http.Request) {
//     c := make(chan []byte)
//     go getGames(w, r, c)
//     respond(w, c)
// }

// func UserShow(w http.ResponseWriter, r *http.Request) {
//     c := make(chan []byte)
//     if db.VerifySessionToken(r) {
//     	go getUser(w, r, c)
//     	respond(w, c)
//     }
// }

// func tttState(w http.ResponseWriter, r *http.Request) {
//     c := make(chan []byte)
//     if db.VerifySessionToken(r) {
//         ttt.Move(w, r, c)
//         respond(w, c)
//     } else {
//         c <- []byte("no bueno")
//         respond(w, c)
//     }
// }

// func hangmanState(w http.ResponseWriter, r *http.Request) {
//     c := make(chan []byte)
//     if db.VerifySessionToken(r) {
//         // go ttt.Move(w, r, c)
//         respond(w, c)
//     }
// }

// func wwufState(w http.ResponseWriter, r *http.Request) {
//     c := make(chan []byte)
//     if db.VerifySessionToken(r) {
//         // go ttt.Move(w, r, c)
//         respond(w, c)
//     }
// }
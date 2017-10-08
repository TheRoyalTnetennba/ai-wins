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
        "Logout",
        "POST",
        "/logout",
        Logout,
    },
	Route{
		"GameRoute",
		"POST",
		"/game/{resource}",
		GameRoute,
	},
}

// ------------ ADD SESSION VALIDATION PRE-CHECK TO GAMEROUTE ------------

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

func Logout(w http.ResponseWriter, r *http.Request) {
    u := User{}
    json.NewDecoder(r.Body).Decode(&u)
    processLogout(u)
    u = User{}
    respond(w, r, 200, u)
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


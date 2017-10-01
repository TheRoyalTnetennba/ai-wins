package main

import ( 
    "net/http"
)

func respond(w http.ResponseWriter, c chan []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.Write(<-c)
}

// PUBLIC GETS

func GameData(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go pubGet(w, r, c)
    respond(w, c)
}

func Auth(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go pubGet(w, r, c)
    respond(w, c)
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go pubGet(w, r, c)
    respond(w, c)
}

// SECRET GETS
func GetUser(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go secGet(w, r, c)
    respond(w, c)
}

// SECRET POSTS

func GetMove(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go secPost(w, r, c)
    respond(w, c)
}
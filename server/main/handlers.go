package main

import ( 
    "encoding/json"
    "github.com/TheRoyalTnetennba/ai-wins/server/utils"
)

// PUBLIC GETS

func GameData(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go pubGet(w, r, c)
    w.Write(<-c)
}

func GamesIndex(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go pubGet(w, r, c)
    w.Write(<-c)
}

func Auth(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go pubGet(w, r, c)
    w.Write(<-c)
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go pubGet(w, r, c)
    w.Write(<-c)
}

// SECRET GETS

// SECRET POSTS

func GetMove(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go postFunc(w, r, c)
    w.Write(<-c)
}
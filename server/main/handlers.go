package main

import ( 
    "net/http"
    "fmt"
    // "github.com/TheRoyalTnetennba/ai-wins/server/utils"
)

// PUBLIC GETS

func GameData(w http.ResponseWriter, r *http.Request) {
    fmt.Println("begin route GameData")
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

// SECRET POSTS

func GetMove(w http.ResponseWriter, r *http.Request) {
    c := make(chan []byte)
    go secPost(w, r, c)
    respond(w, c)
}

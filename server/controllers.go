package main

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"
    "golang.org/x/oauth2"
    "github.com/gorilla/mux"
)

func pubGet(w http.ResponseWriter, r *http.Request, c chan []byte) {
    resource := resourceID(r)
    _, resModel := getRestModels(resource, "")
    var payload []byte
    payload, err := json.Marshal(&resModel)
    if err != nil {
        fmt.Println(err)
    }
    c <- payload
}

func secGet(w http.ResponseWriter, r *http.Request, c chan []byte) {
    resource := resourceID(r)
    token := validToken(r)
    if len(token) < 1 {
        fmt.Println("no token")
    }
    _, resModel := getRestModels(resource, token)
    payload, err := json.Marshal(resModel)
    if err != nil {
        fmt.Println(err)
    }
    c <- payload
}

func secPost(w http.ResponseWriter, r *http.Request, c chan []byte) {
    go respond(w, c)
    vars := mux.Vars(r)
    var payload []byte
    err := json.Unmarshal(payload, vars)
    if err != nil {
        return
    }
    c <- payload
}

func resourceID(r *http.Request) string {
    vars := mux.Vars(r)
    resource := vars["resource"]
    return resource
}

func getRestModels(resource string, token string) (interface{}, interface{}) {
    switch resource {
    case "games":
        return nil, getAllGames()
    case "user":
        return nil, getUserBySessionToken(token)
    }
    return nil, nil
}

func validToken(r *http.Request) string {
    session, err := Store.Get(r, "ai-wins")
    if err != nil {
        fmt.Println(err)
    }
    token := oauth2.Token{
        AccessToken: session.Values["AccessToken"].(string),
        TokenType: session.Values["TokenType"].(string),
        RefreshToken: session.Values["RefreshToken"].(string),
        Expiry: session.Values["Expiry"].(time.Time),
    }
    if token.Expiry.Unix() > time.Now().Unix() {
        return token.AccessToken
    }
    return ""
}
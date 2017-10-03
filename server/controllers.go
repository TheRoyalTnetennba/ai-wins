package main

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"
    "golang.org/x/oauth2"
    "github.com/TheRoyalTnetennba/ai-wins/server/packages/db"

)

func getGames(w http.ResponseWriter, r *http.Request, c chan []byte) {
    games := db.GetAllGames()
    payload, err := json.Marshal(&games)
    if err != nil {
        fmt.Println(err)
    }
    c <- payload
}

func getUser(w http.ResponseWriter, r *http.Request, c chan []byte) {
    user := db.GetUser(r)
    payload, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
    }
    c <- payload
}

func validToken(r *http.Request) string {
    session, err := db.Store.Get(r, "ai-wins")
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

func respond(w http.ResponseWriter, c chan []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.Write(<-c)
}

func problem(w http.ResponseWriter, c chan[]byte, error string, code int) {
    close(c)
    http.Error(w, error, code)
}

package db

import (
    "fmt"
    "time"
    "net/http"
)

func GetUser(r *http.Request) *User {
    token := GetToken(r)
    user := getUserBySessionToken(token)
    return user
}

func Logout(r *http.Request, w http.ResponseWriter) {
    token := GetToken(r)
    deleteLocalToken(r, w)
    deleteRemoteToken(token)
}

func VerifySessionToken(r *http.Request) bool {
    token := GetToken(r)
    expir := GetExpiry(r)
    fmt.Println(expir)
    if expir.Unix() < time.Now().Unix() || len(token) < 1 {
        return false
    }
    return true
}

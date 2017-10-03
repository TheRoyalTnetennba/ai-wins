package db

import (
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

func VerifySessionToken(r *http.Request, w http.ResponseWriter) bool {
    token := GetToken(r)
    expir := GetExpiry(r)
    if expir.Unix() < time.Now().Unix() {
        return false
    }
    return true
}

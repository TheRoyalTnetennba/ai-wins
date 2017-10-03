package db

import (
    "fmt"
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


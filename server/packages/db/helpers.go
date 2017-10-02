package db

import (
    "fmt"
    "net/http"
)

func GetUser(r *http.Request) *User {
    fmt.Println("getting user")
    token := GetToken(r)
    fmt.Println(token)
    fmt.Println("that was token")
    user := getUserBySessionToken(token)
    return user
}

func Logout(r *http.Request, w http.ResponseWriter) {
    token := GetToken(r)
    deleteLocalToken(r, w)
    deleteRemoteToken(token)
}


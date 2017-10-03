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
    fmt.Println("begin token check")
    token := GetToken(r)
    fmt.Println("got token, it is")
    fmt.Println("token")
    fmt.Println("begin get expir")
    expir := GetExpiry(r)
    fmt.Println(expir)
    fmt.Println("that was expiry")
    fmt.Println(expir)
    fmt.Println("begin time check")
    if expir.Unix() < time.Now().Unix() || len(token) < 1 {
        fmt.Println("finish time check it was late")
        return false
    }
    fmt.Println("finish time check it was fine")
    return true
}

package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func pubGet(w http.ResponseWriter, r *http.Request, c chan []byte) {
    resource := resourceID(r)
    _, resModel := getRestModels(resource)
    var payload []byte
    payload, err := json.Marshal(&resModel)
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

func getRestModels(resource string) (interface{}, interface{}) {
    switch resource {
    case "games":
        return nil, getAllGames()
    }
    return nil, nil
}

func respond(w http.ResponseWriter, c chan []byte) {
    fmt.Println("begin respond")
    w.Header().Set("Content-Type", "application/json")
    w.Write(<-c)
}
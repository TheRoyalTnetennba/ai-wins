package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/TheRoyalTnetennba/go-games/ttt"
)

func tttMove(w http.ResponseWriter, r *http.Request, c chan []byte) {
    var marker string
    old := getTTTState(validToken(r))
    current := TTTState{}
    err := json.NewDecoder(r.Body).Decode(current)
    if err != nil || len(current.Marker) < 1 {
        tttSend(w, r, c, old)
    } else if len(old.Marker) > 0 && len(current.Marker) > 0 && old.Marker != current.Marker {
        tttSend(w, r, c, old)
    } else if current.Marker == "x" && matrixIsEmpty(current.Board) {
        old.Marker = current.Marker
        go updateTTTState(old)
        tttSend(w, r, c, old)
    } else if ttt.WhoseTurn(current.Board) == current.Marker {
        tttSend(w, r, c, &current)
    } else if !ttt.IsValidMove(old.Board, current.Board, current.Marker) {
        tttSend(w, r, c, old)
    } else {
        if current.Marker == "o" {
            marker = "x"
        } else {
            marker = "o"
        }
        pos := ttt.GetBestMove(current.Board, marker)
        current.Board[pos[0]][pos[1]] = marker
        go updateTTTState(&current)
        tttSend(w, r, c, &current)
    }
}

func tttSend(w http.ResponseWriter, r *http.Request, c chan []byte, current *TTTState) {
    payload, err := json.Marshal(current)
    if err != nil {
        fmt.Println(err)
    }
    c <- payload
}
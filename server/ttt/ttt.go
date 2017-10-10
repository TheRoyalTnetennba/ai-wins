package main

import (
// "fmt"
)

func processTTT(u User) User {
    o := getUserBySessionToken(u)
    n := newTTTState(o, u)
    go updateUser(n)
    return n
}

func newTTTState(o User, u User) User {
    if !tttValid(o.TTT, u.TTT) {
        return o
    }
    pMarker, aiMarker := tttGetMarkers(u.TTT)
    board := tttGetBestMove(getMatrixFromMap(u.TTT.Board), aiMarker)
    u.TTT.Board = getMapFromMatrix(board)
    switch u.TTT.Result = tttWhoWon(board); u.TTT.Result {
    case "tie":
        u.Tied += 1
        game := getGameFromSlug("tic-tac-toe")
        game.Tied += 1
        u.TTT.Result = "tie"
        go updateGame(game)
    case aiMarker:
        u.Lost += 1
        game := getGameFromSlug("tic-tac-toe")
        game.Won += 1
        u.TTT.Result = "ai"
        go updateGame(game)
    case pMarker:
        u.Won += 1
        game := getGameFromSlug("tic-tac-toe")
        game.Lost += 1
        u.TTT.Result = "player"
        go updateGame(game)
    }
    return u
}

func tttGetMarkers(t TTTState) (string, string) {
    if t.Marker == "o" {
        return "o","x"
    }
    return "x","o"
}

func tttValid(o TTTState, u TTTState) bool {
    return true
}

type TTTState struct {
    Marker string `json:"role,omitempty"`
    Board map[string][]string `json:"board,omitempty"`
    Result string `json:"result,omitempty"`
}
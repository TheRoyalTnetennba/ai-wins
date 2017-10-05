package main

import (
// "fmt"
)

func processTTT(u User) User {
    o := getUserByUID(u.UID)
    n := newTTTState(o, u)
    go updateUser(n)
    return n
}

func newTTTState(o User, u User) User {
    if !isValid(o.TTT, u.TTT) {
        return o
    }
    pMarker, aiMarker := tttGetMarkers(u.TTT)
    board := tttGetBestMove(getMatrixFromMap(u.TTT.Board), aiMarker, aiMarker)
    winner := tttWhoWon(board)
    if winner == "pending" {

    }
    switch winner := tttWhoWon(board); winner {
    case "pending":
        u.TTT.Board = getMapFromMatrix(board)
    case "tie":
        u.TTT.Board = genBoard(3, 3)
        u.Tied += 1
        game := getGameFromSlug("tic-tac-toe")
        game.Tied += 1
        go updateGame(game)
    case aiMarker:
        u.TTT.Board = genBoard(3, 3)
        u.Lost += 1
        game := getGameFromSlug("tic-tac-toe")
        game.Won += 1
        go updateGame(game)
    case pMarker:
        u.TTT.Board = genBoard(3, 3)
        u.Won += 1
        game := getGameFromSlug("tic-tac-toe")
        game.Lost += 1
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

func isValid(o TTTState, u TTTState) bool {
    return true
}
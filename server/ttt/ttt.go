package main

import (

)

func newTTTState(o TTTState, u TTTState) TTTState {
    if !isValid(o, u) {
        return o
    }
    pMarker, aiMarker := tttGetMarkers(u)
    board := getMatrixFromMap(u.Board)
    pos := tttGetBestMove(board, pMarker, aiMarker)
    board[pos[0]][pos[1]] = aiMarker
    u.Board = getMapFromMatrix(board)
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

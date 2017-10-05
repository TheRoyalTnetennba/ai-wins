package main

import (
	"fmt"
	"sort"
)

type tttGameNode struct {
	Board     [][]string
	NextGames [9]*tttGameNode
	Result    int
}

func tttNextFree(arr [9]*tttGameNode) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] == nil })
}

func tttGenNodes(game *tttGameNode, marker string, layer int, aiMarker string) int {
	winner := tttWhoWon(game.Board)
	if winner != "pending" {
		if winner == "tie" {
			game.Result = 0
		} else if winner == aiMarker {
			game.Result = 1 * layer
		} else {
			game.Result = -1 * layer
		}
		return game.Result
	}
	for i := 0; i < len(game.Board); i++ {
		for j := 0; j < len(game.Board[i]); j++ {
			if game.Board[i][j] == "" {
				childBoard := copyMatrix(game.Board)
				childBoard[i][j] = marker
				var childNextGames [9]*tttGameNode
				childResult := 0
				child := tttGameNode{childBoard, childNextGames, childResult}
				game.NextGames[tttNextFree(game.NextGames)] = &child
			}
		}
	}
	if marker == "x" {
		marker = "o"
	} else {
		marker = "x"
	}
	for _, child := range game.NextGames {
		if child != nil {
			game.Result += tttGenNodes(child, marker, layer/10, aiMarker)
		}
	}
	return game.Result
}

func tttCheckTriplet(triplet []string) string {
	if triplet[0] == triplet[1] && triplet[1] == triplet[2] {
		return triplet[0]
	}
	for i := 0; i < len(triplet); i++ {
		if triplet[i] == "" {
			return "Incomplete"
		}
	}
	return "Contested"
}

func tttWin(tripletResponse string) bool {
	return tripletResponse == "x" || tripletResponse == "o"
}

func tttWhoWon(board [][]string) string {
	numIncompletes := 0
	for row := 0; row < len(board); row++ {
		rowRes := tttCheckTriplet(board[row])
		if tttWin(rowRes) {
			return rowRes
		} else if rowRes == "Incomplete" {
			numIncompletes++
		}
		var colRow []string
		for col := 0; col < len(board[row]); col++ {
			colRow = append(colRow, board[col][row])
		}
		colRes := tttCheckTriplet(colRow)
		if tttWin(colRes) {
			return colRes
		} else if colRes == "Incomplete" {
			numIncompletes++
		}
	}
	var upRight []string
	var upLeft []string
	for i := 0; i < len(board); i++ {
		upRight = append(upRight, board[i][len(board[i])-(1+i)])
		upLeft = append(upLeft, board[i][i])
	}
	upRightRes, upLeftRes := tttCheckTriplet(upRight), tttCheckTriplet(upLeft)
	if tttWin(upLeftRes) {
		return upLeftRes
	} else if tttWin(upRightRes) {
		return upRightRes
	} else if numIncompletes > 0 {
		return "pending"
	}
	return "tie"
}

func tttIsValidMove(old [][]string, current [][]string, marker string) bool {
	changes := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if len(old[i][j]) > 0 && old[i][j] != current[i][j] {
				return false
			} else if old[i][j] != current[i][j] && current[i][j] != marker {
				return false
			} else if old[i][j] != current[i][j] {
				changes += 1
			}
		}
	}
	return changes == 1
}

func tttWhoseTurn(current [][]string) string {
	x, o := 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if current[i][j] == "x" {
				x += 1
			} else if current[i][j] == "o" {
				o += 1
			}
		}
	}
	if x > o {
		return "o"
	}
	return "x"
}

func tttGetBestMove(board [][]string, marker string, aiMarker string) [][]string {
	var children [9]*tttGameNode
	res := 0
	game := tttGameNode{board, children, res}
	if tttWhoWon(game.Board) != "pending" {
		return board
	}
	tttGenNodes(&game, marker, 1000000000, aiMarker)
	max := game.NextGames[0]
	for _, child := range game.NextGames {
		if child != nil && child.Result > max.Result {
			fmt.Println(child)
			max = child
		}
	}
	return max.Board
}

package tttAI

import (
	"fmt"

	"github.com/TheRoyalTnetennba/ai-wins/server/utils"
)

var (
	aiMarker string
	games    []gameNode
)

type gameNode struct {
	Board     [][]string
	NextGames []gameNode
	Result    int
}

func genNodes(game gameNode, marker string) int {
	winner := whoWon(game)
	if winner != "pending" {
		if winner == "tie" {
			game.Result = 0
		} else if winner == aiMarker {
			game.Result = 1
		} else {
			game.Result = -1
		}
		return game.Result
	}
	for i := 0; i < len(game.Board); i++ {
		for j := 0; j < len(game.Board[i]); j++ {
			if game.Board[i][j] == "" {
				fmt.Println("here is an empty one")
				childBoard := utils.CopyMatrix(game.Board)
				childBoard[i][j] = marker
				var childNextGames []gameNode
				childResult := 0
				child := gameNode{childBoard, childNextGames, childResult}
				parent := &game
				parent.NextGames = append(parent.NextGames, child)
			}
		}
	}
	if marker == "x" {
		marker = "o"
	} else {
		marker = "x"
	}
	for _, child := range game.NextGames {
		child.Result += genNodes(child, marker)
	}
	return game.Result
}

func checkTriplet(triplet []string) string {
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

func win(tripletResponse string) bool {
	return tripletResponse == "x" || tripletResponse == "o"
}

func whoWon(game gameNode) string {
	board := game.Board
	numIncompletes := 0
	for row := 0; row < len(board); row++ {
		rowRes := checkTriplet(board[row])
		if win(rowRes) {
			return rowRes
		} else if rowRes == "Incomplete" {
			numIncompletes++
		}
		var colRow []string
		for col := 0; col < len(board[row]); col++ {
			colRow = append(colRow, board[col][row])
		}
		colRes := checkTriplet(colRow)
		if win(colRes) {
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
	upRightRes, upLeftRes := checkTriplet(upRight), checkTriplet(upLeft)
	if win(upLeftRes) {
		return upLeftRes
	} else if win(upRightRes) {
		return upRightRes
	} else if numIncompletes > 0 {
		return "pending"
	}
	return "tie"
}

func DumMove(board [][]string) []int {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if len(board[i][j]) == 0 {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func boardDifference(parent [][]string, child [][]string) []int {
	for i := 0; i < len(parent); i++ {
		for j := 0; j < len(parent[i]); j++ {
			if parent[i][j] != child[i][j] {
				return []int{i, j}
			}
		}
	}
	return DumMove(parent)
}

func GetAIMove(board [][]string, marker string) []int {
	aiMarker = marker
	var children []gameNode
	res := 0
	game := gameNode{board, children, res}
	genNodes(game, marker)
	fmt.Println(game)
	max := game.NextGames[0]
	for _, child := range game.NextGames {
		if child.Result > max.Result {
			max = child
		}
	}
	return boardDifference(board, max.Board)
}

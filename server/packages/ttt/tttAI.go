package ttt

import (
	"sort"
)

var (
	aiMarker string
)

type gameNode struct {
	Board     [][]string
	NextGames [9]*gameNode
	Result    int
}

func nextFree(arr [9]*gameNode) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] == nil })
}

func genNodes(game *gameNode, marker string, layer int) int {
	winner := whoWon(game)
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
				var childNextGames [9]*gameNode
				childResult := 0
				child := gameNode{childBoard, childNextGames, childResult}
				game.NextGames[nextFree(game.NextGames)] = &child
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
			game.Result += genNodes(child, marker, layer/10)
		}
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

func whoWon(game *gameNode) string {
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

func copyMatrix(orig [][]string) [][]string {
	var newMatrix [][]string
	for i := 0; i < len(orig); i++ {
		var row []string
		for j := 0; j < len(orig[i]); j++ {
			row = append(row, orig[i][j])
		}
		newMatrix = append(newMatrix, row)
	}
	return newMatrix
}

func IsValidMove(old [][]string, current [][]string, marker string) bool {
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

func WhoseTurn(current [][]string) string {
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

func GetBestMove(board [][]string, marker string) []int {
	aiMarker = marker
	var children [9]*gameNode
	res := 0
	game := gameNode{board, children, res}
	if whoWon(&game) != "pending" {
		return []int{0, 0}
	}
	genNodes(&game, marker, 1000000000)
	max := game.NextGames[0]
	for _, child := range game.NextGames {
		if child != nil && child.Result > max.Result {
			max = child
		}
	}
	return boardDifference(board, max.Board)
}
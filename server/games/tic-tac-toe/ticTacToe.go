package tttAI

// func checkTriplet(triplet []string) string {
// 	if triplet[0] == triplet[1] && triplet[1] == triplet[2] {
// 		return triplet[0]
// 	}
// 	for i := 0; i < len(triplet); i++ {
// 		if triplet[i] == "" {
// 			return "Incomplete"
// 		}
// 	}
// 	return "Contested"
// }
// func win(tripletResponse string) bool {
// 	return tripletResponse == "x" || tripletResponse == "o"
// }

// func isGameOver(board [][]string) bool {
// 	numIncompletes := 0
// 	for row := 0; row < len(board); row++ {
// 		if win(checkTriplet(board[row])) {
// 			return true
// 		}
// 		var colRow []string
// 		for col := 0; col < len(board[row]); col++ {
// 			colRow = append(colRow, board[col][row])
// 		}
// 		carlRove := checkTriplet(colRow)
// 		if win(carlRove) {
// 			return true
// 		} else if carlRove == "Incomplete" {
// 			numIncompletes++
// 		}
// 		if win(checkTriplet(colRow)) {
// 			return true
// 		}
// 	}
// 	var upRight []string
// 	var upLeft []string
// 	for i := 0; i < len(board); i++ {
// 		upRight = append(upRight, board[i][len(board[i])-i])
// 		upLeft = append(upLeft, board[i][i])
// 	}
// 	if win(checkTriplet(upRight)) || win(checkTriplet(upLeft)) {
// 		return true
// 	}
// 	return numIncompletes == 0
// }

func GetAIMove(board [][]string, marker string) []int {
	if len(board[1][1]) == 0 {
		return []int{1, 1}
	} else {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if len(board[i][j]) == 0 {
					return []int{i, j}
				}
			}
		}
	}
	return []int{0, 0}
}

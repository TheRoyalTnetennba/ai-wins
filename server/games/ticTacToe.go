package main

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

func tttIsGameOver(board [][]string) bool {
	numIncompletes := 0
	for row := 0; row < len(board); row++ {
		if tttWin(tttCheckTriplet(board[row])) {
			return true
		}
		var colRow []string
		for col := 0; col < len(board[row]); col++ {
			colRow = append(colRow, board[col][row])
		}
		carlRove := tttCheckTriplet(colRow)
		if tttWin(carlRove) {
			return true
		} else if carlRove == "Incomplete" {

		}
		if tttWin(tttCheckTriplet(colRow)) {
			return true
		}
	}
	var upRight []string
	var upLeft []string
	for i := 0; i < len(board); i++ {
		upRight = append(upRight, board[i][len(board[i])-i])
		upLeft = append(upLeft, board[i][i])
	}
	if tttWin(tttCheckTriplet(upRight)) || tttWin(tttCheckTriplet(upLeft)) {
		return true
	}
	return numIncompletes == 0
}

// 1.

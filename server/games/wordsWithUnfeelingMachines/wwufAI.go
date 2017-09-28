package wwuf

import (
    "strconv"
)

// !!!!!! NEED TO ACCOUNT FOR WILDCARD CHARACTERS !!!!!!

func scoreLetter(letter string) int {
    letters := []string{"D", "G", "B", "C", "M", "P", "F", "H", "V", "W", "Y", "K", "J", "X", "Q", "Z"}
    points := []int{2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 8, 8, 10, 10}
    for i := 0; i < len(letterScore); i++ {
        if letters[i] == letter {
            return points[i]
        }
    }
    return 1
}

func getMultiplier(pos string) []int {
    wl := []int{1, 1}
    tw := []string{"0,0", "0,7", "0,14", "7,0", "7,14", "14,0", "14,7", "14,14"}
    dl := []string{"0,3", "0,11", "7,3", "7,11", "14,3", "3,0", "11,0", "3,14", "11,14", "14,11", "3,7", "11,7",
                   "2,6", "2,8", "12,6", "12,8", "6,2", "6,6", "6,8", "6,12", "8,2", "8,6", "8,8", "8,12"}
    tl := []string{"1,5", "1,9", "5,1", "5,5", "5,9", "5,13", "9,1", "9,5", "9,9", "9,13", "13,5", "13,9"}
    dw := []string{"1,1", "2,2", "3,3", "4,4", "7,7", "10,10", "11,11", "12,12", "13,13", "13,1", "12,2", "11,3",
                   "10,4", "4,10", "3,11", "2,12", "1,13"}
    for i := 0; i < len(tw); i++ {
        if tw[i] == pos {
            wl[0] *= 3
        }
    }
    for i := 0; i < len(dl); i++ {
        if dl[i] == pos {
            wl[1] *= 2
        }
    }
    for i := 0; i < len(tl); i++ {
        if tl[i] == pos {
            wl[1] *= 3
        }
    }
    for i := 0; i < len(dw); i++ {
        if dw[i] == pos {
            wl[1] *= 2
        }
    }
    return wl
}

func ScoreMove(newBoard [][]string, oldBoard [][]string) (int, [][]int) {
    var positions [][]int
    multiplier, points := 1, 0
    for i := 0; i < len(oldBoard); i++ {
        for j := 0; j < len(oldBoard[i]); j++ {
            if newBoard[i][j] != oldBoard[i][j] {
                position := []int{i, j}
                pos := strconv.Itoa(i) + "," + strconv.Itoa(j)
                mult := getMultiplier(pos)
                multiplier *= mult[0]
                points += mult[1] * scoreLetter(newBoard[i][j])
                positions = append(positions, position)
            }
        }
    }
    points *= multiplier
    deltaX := positions[0][1] - positions[1][1]
    deltaY := positions[0][0] - positions[0][1]
    if deltaX == 0 {
        for i := 1; i < len(positions); i++ {
            if positions[i - 1][1] + 1 < positions[i][1] {

            }
        }
    }
    return points * multiplier
}


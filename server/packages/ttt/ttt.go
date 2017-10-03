package ttt

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/TheRoyalTnetennba/ai-wins/server/packages/db"
    "github.com/TheRoyalTnetennba/ai-wins/server/packages/utils"
)

func Move(w http.ResponseWriter, r *http.Request, c chan []byte) {
    fmt.Println("starting move")
    user, aiMarker := db.GetUser(r), "x"
    old := db.GetTTTState(user)
    current := &db.TTTState{}
    err := json.NewDecoder(r.Body).Decode(current)
    if err != nil {
        fmt.Println("error reading body")
        fmt.Println(err)
    }
    patchCurrent(old, current)
    if Valid(old, current) {
        current.User = old.User
        if current.Marker == "x" {
            aiMarker = "o"
        }
        pos := GetBestMove(current.Board, aiMarker)
        current.Board[pos[0]][pos[1]] = aiMarker
        db.UpdateTTTState(current)
        tttSend(c, current)
    }
}

func patchCurrent(old *db.TTTState, current *db.TTTState) {
    if len(current.Marker) == 0 {
        current.Marker = old.Marker
    }
    current.User = old.User
    current.Game = old.Game
    current.Started = old.Started
    current.Key = old.Key
    if len(current.Board) == 0 {
        current.Board = utils.NewMatrix(3)
    }
}

func Valid(old *db.TTTState, current *db.TTTState) bool {
    return true
}

// func isValid(old *TTTState, current TTTState) bool {
//     deltaBoard := len(boardChanges(old.Board, current.Board)) == 1
//     noCheat := noOverrides(old.Board, current.Board)


//     return deltaBoard && noCheat
// }

// func isOver() bool string {

//     return false
// }

// func noOverrides(old [][]string, current [][]string) bool {
//     for i := 0; i < len(old); i++ {
//         for j := 0; j < len(old[i]); j++ {
//             if len(old[i][j]) > 0 && old[i][j] != current[i][j] {
//                 return false
//             }
//         }
//     }
//     if len(old.Marker) > 0 && old.Marker != current.Marker {
//         return false
//     }
//     return true
// }

// func boardChanges(old [][]string, current [][]string) [][]int {
//     var diffs [][]int
//     for i := 0; i < len(old); i++ {
//         for j := 0; j < len(old[i]); j++ {
//             if old[i][j] != current[i][j] {
//                 diff := []int{i,j}
//                 diffs = append(diffs, diff)
//             }
//         }
//     }
//     return diffs
// }







// func tttMove(w http.ResponseWriter, r *http.Request, c chan []byte) {
//     var marker string
//     token := validToken(r)
//     old := getTTTState(token)
//     current := TTTState{}
//     err := json.NewDecoder(r.Body).Decode(current)
//     if err != nil || len(current.Marker) < 1 {
//         tttSend(w, r, c, old)
//     } else if len(old.Marker) > 0 && len(current.Marker) > 0 && old.Marker != current.Marker {
//         tttSend(w, r, c, old)
//     } else if current.Marker == "x" && matrixIsEmpty(current.Board) {
//         old.Marker = current.Marker
//         go updateTTTState(old)
//         tttSend(w, r, c, old)
//     } else if ttt.WhoseTurn(current.Board) == current.Marker {
//         tttSend(w, r, c, &current)
//     } else if !ttt.IsValidMove(old.Board, current.Board, current.Marker) {
//         tttSend(w, r, c, old)
//     } else {
//         if current.Marker == "o" {
//             marker = "x"
//         } else {
//             marker = "o"
//         }
//         pos := ttt.GetBestMove(current.Board, marker)
//         current.Board[pos[0]][pos[1]] = marker
//         go updateTTTState(&current)
//         tttSend(w, r, c, &current)
//     }
// }

// func attachUserToTTTState(token string, current *TTTState) {
//     user := getUserBySessionToken(token)
//     current.User = user.Key
// }

func tttSend(c chan []byte, current *db.TTTState) {
    payload, err := json.Marshal(current)
    if err != nil {
        fmt.Println(err)
    }
    c <- payload
}
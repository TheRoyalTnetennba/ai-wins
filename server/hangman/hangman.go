package main

import (
// "fmt"
    "time"
)

type HangmanState struct {
    UserGuess bool `json:"userGuess"`
    Started time.Time `json:"started"`
    Board []string `json:"board"`
    Word string `json:"board"`
    Misses []string `json:"misses"`
    Result string `json:"result"`
}



// func processHangman(u User) User {
//     o := getUserByUID(u.UID)
//     n := newTTTState(o, u)
//     go updateUser(n)
//     return n
// }

// func hangmanValid(o HangmanState, u HangmanState) bool {
//     return true
// }

// func (u HangmanState) PossibleWords() []string {

// }

// func (w SowpodEntry) Contains(letter string) bool {
//     return arrayIncludes(w.WordArr, letter)
// }

// func (w SowpodEntry) Matches(letters []string) bool {
//     if len(letters) != len(w.WordArr) {
//         return false
//     }
//     for let i := 0; i < len(letters); i++ {
//         if letters[i] == "_" {
//             continue
//         }
//         if letters[i] != w.WordArr[i] {
//             return false
//         }
//     }
//     return true
// }



// func (u HangmanState) AI() HangmanState {
//     return HangmanState{}
// }

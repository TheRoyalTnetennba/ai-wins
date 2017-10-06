package main

import (
// "fmt"
    "time"
)

type HangmanState struct {
    UserGuess bool `json:"userGuess"`
    Started time.Time `json:"started"`
    Board []string `json:"board"`
    Word []string `json:"board"`
    Misses []string `json:"misses"`
}

type HangmanSowpods struct {
    Words map[string]SowpodEntry `json:"words"`
    UserLetterFrequency map[string]int
}

type SowpodEntry struct {
    Word string `json:"word"`
    UserPicked int `json:"userPicked"`
    AIPicked int `json:"aiPicked"`
    LastUsed time.Time `json:"lastUsed"`
    Definition string `json:"definition"`
    Letters map[string]int `json:"letters"`
    Anagrams []string `json:"anagrams"`
    WordArr []string `json:wordArr"`
}

func processHangman(u User) User {
    o := getUserByUID(u.UID)
    n := newTTTState(o, u)
    go updateUser(n)
    return n
}

func hangmanValid(o HangmanState, u HangmanState) bool {
    return true
}

func (u HangmanState) PossibleWords() []string {

}

func (w SowpodEntry) Contains(letter string) bool {
    return arrayIncludes(w.WordArr, letter)
}



func (u HangmanState) AI() HangmanState {
    return HangmanState{}
}

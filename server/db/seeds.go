package main

import (
    // "sort"
    "strings"
    "time"
    "fmt"
)

type SowpodEntry struct {
    Word string `json:"word"`
    UserPicked int `json:"userPicked"`
    AIPicked int `json:"aiPicked"`
    LastUsed time.Time `json:"lastUsed"`
    Definition string `json:"definition"`
    Letters map[string]int `json:"letters"`
    Anagrams []string `json:"anagrams"`
    WordArr []string `json:wordArr"`
    Sorted string `json:"string"`
    Length int
}

func seedSowpods() {
    sortagrams := make(map[string][]string)
    words := ReadLines("./server/utils/sowpods.txt")
    for _, word := range words {
        sorted := sortString(word)
        if val, ok := sortagrams[sorted]; ok {
            newVal := append(val, word)
            sortagrams[sorted] = newVal
        } else {
            newVal := []string{word}
            sortagrams[sorted] = newVal
        }
    }
    // entries := make(map[string]SowpodEntry)
    for i := 30000; i < 45000 && i < len(words); i++ {
        word := words[i]
        entry := SowpodEntry{
            Word: word,
            UserPicked: 1,
            AIPicked: 1,
            LastUsed: time.Now(),
            Definition: "",
            Letters: letterFrequency(word),
            Anagrams: sortagrams[sortString(word)],
            WordArr: strings.Split(word, ""),
            Sorted: sortString(word),
        }
        
        go func () {
            _, err := client.Collection("sowpods").Doc(word).Set(ctx, &entry)
            if err != nil {
                fmt.Println(err)
                fmt.Println(word)
                fmt.Println(i)
            }
        }()
    }
}
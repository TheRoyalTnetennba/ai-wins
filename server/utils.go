package main

import (
    "os"
    "fmt"
    "sort"
    "time"
    "bufio"
    "strings"
	"math/rand"
)

var (
	Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	rand.Seed(time.Now().Unix())
}

func readLines(path string) ([]string) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("error reading file")
        return nil
    }
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func getSowpods() (map[string][]string) {
    dict := make(map[string][]string)
    words := readLines("./assets/sowpods.txt")
    for _, word := range words {
        arr := strings.Split(word, "")
        sort.Strings(arr)
        sorted := strings.Join(arr, "")
        if val, ok := dict[sorted]; ok {
            newVal := append(val, word)
            dict[sorted] = newVal
        } else {
            newVal := []string{word}
            dict[sorted] = newVal
        }
    }
    return dict
}

func matrixIsEmpty(board [][]string) bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if len(board[i][j]) > 0 {
                return false
            }
        }
    }
    return true
}

func newMatrix(n int) [][]string {
    var newMatrix [][]string
    for i := 0; i < n; i++ {
        var row []string
        for j := 0; j < n; j++ {
            row = append(row, "")
        }
        newMatrix = append(newMatrix, row)
    }
    return newMatrix
}

func randNum(upperLimit int) int {
	return rand.Intn(upperLimit)
}

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = Letters[rand.Intn(len(Letters))]
    }
    return string(b)
}
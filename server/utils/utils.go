package main

import (
    "os"
    "fmt"
    "sort"
    "time"
    "bufio"
    "strings"
    "math/rand"
    "strconv"
)

var (
    Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    Dict map[string][]string = GetSowpods()
)

func init() {
    rand.Seed(time.Now().Unix())
}

func ReadLines(path string) ([]string) {
    fmt.Println(path)
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

func GetSowpods() (map[string][]string) {
    dict := make(map[string][]string)
    words := ReadLines("./server/utils/sowpods.txt")
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

func genBoard(r int, c int) map[string][]string {
    board := make(map[string][]string)
    for i := 0; i < r; i++ {
        a := strconv.Itoa(i)
        var arr []string
        board[a] = arr
        for j := 0; j < c; j++ {
            board[a] = append(board[a], "")
        }
    }
    return board
}

func getMatrixFromMap(board map[string][]string) [][]string {
    var grid [][]string
    for i := 0; true; i++ {
        a := strconv.Itoa(i)
        if val, ok := board[a]; ok {
            grid = append(grid, val)
        } else {
            return grid
        }
    }
    return grid
}

func getMapFromMatrix(grid [][]string) map[string][]string {
    board := make(map[string][]string)
    for i := 0; i < len(grid); i++ {
        a := strconv.Itoa(i)
        var arr []string
        for j := 0; j < len(grid[i]); j++ {
            arr = append(arr, grid[i][j])
        }
        board[a] = arr
    }
    return board
}

func arrayIncludes(arr []string, letter string) bool {
    for i := 0; i < len(arr); i++ {
        if arr[i] == letter {
            return true
        }
    }
    return false
}

func arrayIncludesAny(arr []string, letters []string) bool {
    for i := 0; i < len(letters); i++ {
        if arrayIncludes(arr, letters[i]) {
            return true
        }
    }
    return true
}
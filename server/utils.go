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

// type User struct {
//     Image string
//     Joined time.Time
//     Lost int
//     Tied int
//     UserName string
//     Won int
//     Email string
//     Token string
// }

// type GoogleUser struct {
//     Email string `json:"email,omitempty"`
//     LastName string `json:"family_name,omitempty"`
//     Gender string `json:"gender,omitempty"`
//     FirstName string `json:"given_name,omitempty"`
//     Id string `json:"id,omitempty"`
//     Link string `json:"link,omitempty"`
//     Locale string `json:"locale,omitempty"`
//     Name string `json:"name,omitempty"`
//     Picture string `json:"picture,omitempty"`
//     VerifiedEmail bool `json:"verified_email,omitempty"`
// }

// func normalizeOUser(oUser interface{}) User {
//     switch oUser.(type) {
//     case GoogleUser:
//         return normalizeGoogleUser(oUser.(GoogleUser))
//     }
//     return User{}
// }

// func normalizeGoogleUser(gUser GoogleUser) User {
//     user := User{
//         Image: gUser.Picture,
        
//     }
//     return User{}
// }
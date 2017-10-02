package wwuf

import (
    // "fmt"
    // "net/http"

    // simplejson "github.com/bitly/go-simplejson"
    // "github.com/gorilla/mux"
)

func NewLetterSet() []string {
    var letters []string
    for i := 0; i < 12; i++ {
        letters = append(letters, "E")
    }
    for i := 0; i < 9; i++ {
        letters = append(letters, "A")
        letters = append(letters, "I")
    }
    for i := 0; i < 8; i++ {
        letters = append(letters, "O")
    }
    for i := 0; i < 6; i++ {
        letters = append(letters, "N")
        letters = append(letters, "R")
        letters = append(letters, "T")
    }
    for i := 0; i < 4; i++ {
        letters = append(letters, "L")
        letters = append(letters, "S")
        letters = append(letters, "U")
        letters = append(letters, "D")
    }
    for i := 0; i < 3; i++ {
        letters = append(letters, "G")
    }
    for i := 0; i < 2; i++ {
        letters = append(letters, "B")
        letters = append(letters, "C")
        letters = append(letters, "M")
        letters = append(letters, "P")
        letters = append(letters, "F")
        letters = append(letters, "H")
        letters = append(letters, "V")
        letters = append(letters, "W")
        letters = append(letters, "Y")
        letters = append(letters, "-")
    }
    letters = append(letters, "K")
    letters = append(letters, "J")
    letters = append(letters, "X")
    letters = append(letters, "Q")
    letters = append(letters, "Z")
    return letters
}
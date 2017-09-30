package main

type GoogleUser struct {
    Email string
    Family_name string
    Gender string
    Given_name string
    Id string
    Link string
    Locale string
    Name string
    Picture string
    Verified_email bool
}

type TTTMove struct {
    Board [][]string
    Marker string
    Move [] int
}

type HangmanMove struct {
    Board [][]string

}

type Games []*Game
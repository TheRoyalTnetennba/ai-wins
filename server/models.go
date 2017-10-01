package main

import (
    "time"
)

type GoogleUser struct {
    Email string `json:"email,omitempty"`
    LastName string `json:"family_name,omitempty"`
    Gender string `json:"gender,omitempty"`
    FirstName string `json:"given_name,omitempty"`
    ID string `json:"id,omitempty"`
    Link string `json:"link,omitempty"`
    Locale string `json:"locale,omitempty"`
    Name string `json:"name,omitempty"`
    Picture string `json:"picture,omitempty"`
    VerifiedEmail bool `json:"verified_email,omitempty"`
}

type TTTMove struct {
    Board [][]string
    Marker string
    Move [] int
}

type HangmanMove struct {
    Board [][]string

}

type AuthUrl struct {
    URL string 
}

type Game struct {
    CanTie bool
    DateAdded time.Time
    Description string
    GamesStarted int
    LastLossDate time.Time
    LastLossUsername string
    Lost int
    Name string
    Slug string
    Tagline string
    Thumbnail string
    Tied int
    Won int
}

type User struct {
    Image string
    Joined time.Time
    Lost int
    Tied int
    UserName string
    Won int
    Email string
    Token string
    OAuthID string
}

type Token struct {
    AccessToken string
    TokenType string
    RefreshToken string
    Expiry time.Time
}

type Games []*Game
type Users []*User
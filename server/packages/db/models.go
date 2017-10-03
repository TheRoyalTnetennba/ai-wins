package db

import (
    "time"
    "cloud.google.com/go/datastore"
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

type AuthUrl struct {
    URL string 
}

type Game struct {
    Key *datastore.Key `datastore:"__key__"`
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
    Username string
    Won int
    Email string
    Token string
    OAuthID string
    Key *datastore.Key `datastore:"__key__"`
}

type TTTState struct {
    User *datastore.Key
    Game *datastore.Key
    Marker string
    Started time.Time
    Board []string
    Grid [][]string `datastore:"-"`
    Key *datastore.Key `datastore:"__key__"`
}

type HangmanState struct {
    User *datastore.Key
    Game *datastore.Key
    UserGuess bool
    Started time.Time
    Board []string
    Word []string
    Misses []string
    Key *datastore.Key `datastore:"__key__"`
}

type WWUFState struct {
    User *datastore.Key
    Game *datastore.Key
    Started time.Time
    Board []string
    Grid [][]string `datastore:"-"`
    UserLetters []string
    AILetters []string
    B1 string
    B2 string
    RemainingLetters []string
    Key *datastore.Key `datastore:"__key__"`
}

type Games []*Game
type Users []*User
type TTTStates []*TTTState
type HangmanStates []*HangmanState
type WWUFStates []*WWUFState

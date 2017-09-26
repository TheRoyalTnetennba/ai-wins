package main

import (
    "time"
)

type User struct {
    Image string
    Joined time.Time
    Lost int
    Tied int
    UserName string
    Won int
}

type Game struct {
    CanTie bool
    DateAdded time.Time
    Description string
    GamesStarted int
    Lost int
    Name string
    Slug string
    Tagline string
    Tied int
    Won int
}
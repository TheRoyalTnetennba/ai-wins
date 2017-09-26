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
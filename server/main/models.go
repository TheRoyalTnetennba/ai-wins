package main

import (
    "fmt"
    "log"
    "time"
    "cloud.google.com/go/datastore"
    "golang.org/x/net/context"
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
    GamesStarted int
    Lost int
    Name string
    Slug string
    Tied int
    Won int
}
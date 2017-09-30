package main

import (
    "time"
    "cloud.google.com/go/datastore"
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

type DictWord struct {
    Anagrams []string `datastore:",noindex"`
    Key *datastore.Key `datastore:"__key__"`
}
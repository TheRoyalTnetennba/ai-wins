package main

import (
    // "fmt"
    "time"
    // "strconv"
)

func getAllGames() Games {
    games := Games{}
    iter := client.Collection("games").Documents(ctx)
    for {
        doc, err := iter.Next()
        if err != nil {
            break
        }
        game := Game{}
        doc.DataTo(&game)
        games = append(games, game)
    }
    return games
}

func processGoogleLogin(g GoogleUser) User {
    u := getUserByUID(g.UID)
    if u.UID != g.UID {
        u.Username, u.Email, u.PhoneNumber, u.Image, u.UID, u.Token = g.Name, g.Email, g.PhoneNumber, g.Picture, g.UID, g.Token
        u.Joined = time.Now()
        go addUser(normalizeUser(u))
        return u
    }
    return u
}

func processTTT(u User) User {
    o := getUserByUID(u.UID)
    n := newTTTState(o.TTT, u.TTT)
    u.TTT = n
    go updateUser(u)
    return u
}
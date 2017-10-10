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
    go setSessionToken(u, g.Token)
    return u
}

// !!!!!!!!!!!! THIS IS REALLY STUPID AND DANGEROUS. FIX ASAP !!!!!!!!!!!!
func updateGame(game Game) {
    client.Collection("games").Doc(game.Slug).Set(ctx, &game)
}

func processLogout(u User) {
    user := getUserBySessionToken(u)
    go setSessionToken(user, OToken{})
}
package main

import (
"fmt"
    "cloud.google.com/go/datastore"
)

func getAllGames() Games {
    games := Games{}
    Client.GetAll(Ctx, datastore.NewQuery("Game"), &games)
    return games
}

func getUserBySessionToken(token string) *User {

    users := Users{}

    query := datastore.NewQuery("User")
    Client.GetAll(Ctx, query.Filter("Token=", token), &users)
    fmt.Println(users)
    if len(users) > 0 {
        return users[0]
    }
    return &User{}
}

func getUserByOAuthID(id string, token string) Users {
    users := Users{}
    query := datastore.NewQuery("User")
    keys, _ := Client.GetAll(Ctx, query.Filter("OAuthID=", id), &users)
    if len(users) > 0 {
        users[0].Token = token
        Client.Put(Ctx, keys[0], users[0])
    }
    return users
}

func addNewUser(user *User) error {
    newKey := datastore.IncompleteKey("User", nil)
    _, err := Client.Put(Ctx, newKey, user)
    if err != nil {
        return err
    }
    return nil
}
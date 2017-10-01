package main

import (
    "net/http"
    "golang.org/x/oauth2"
    "cloud.google.com/go/datastore"
)

func getAllGames() Games {
    games := Games{}
    Client.GetAll(Ctx, datastore.NewQuery("Game"), &games)
    return games
}

func getUserBySessionToken(r *http.Request) Users {
    session, _ := Store.Get(r, "ai-wins")
    inter, ok := session.Values["sessionToken"]
    token := inter.(oauth2.Token)
    users := Users{}
    if ok == false {
        return users
    } else if !token.Valid() {
        return users
    }
    query := datastore.NewQuery("User")
    Client.GetAll(Ctx, query.Filter("Token=", token.AccessToken), &users)
    return users
}

func getUserByOAuthID(id string) Users {
    users := Users{}
    query := datastore.NewQuery("User")
    Client.GetAll(Ctx, query.Filter("OAuthID=", id), &users)
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
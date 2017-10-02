package main

import (
    "fmt"
    "time"
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
    if len(users) > 0 {
        fmt.Println(users[0].Key)
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

func deleteOAuthToken(token string) error {
    user := getUserBySessionToken(token)
    if len(user.Key.String()) > 0 {
        user.Token = ""
        Client.Put(Ctx, user.Key, user)
    }
    return nil
}

func incrementWon(token string) error {
    user := getUserBySessionToken(token)
    if len(user.Key.String()) > 0 {
        user.Won = user.Won + 1
        Client.Put(Ctx, user.Key, user)
    }
    return nil
}

func incrementLost(token string) error {
    user := getUserBySessionToken(token)
    if len(user.Key.String()) > 0 {
        user.Lost = user.Lost + 1
        Client.Put(Ctx, user.Key, user)
    }
    return nil
}

func incrementTied(token string) error {
    user := getUserBySessionToken(token)
    if len(user.Key.String()) > 0 {
        user.Tied = user.Tied + 1
        Client.Put(Ctx, user.Key, user)
    }
    return nil
}

func updateUserName(token string, newUsername string) error {
    user := getUserBySessionToken(token)
    if len(user.Key.String()) > 0 {
        user.Username = newUsername
        Client.Put(Ctx, user.Key, user)
    }
    return nil
}

func getTTTState(token string) *TTTState {
    user := getUserBySessionToken(token)
    tttStates := TTTStates{}
    Client.GetAll(Ctx, datastore.NewQuery("TTTState").Filter("User=", user.Key), &tttStates)
    if len(tttStates) > 0 {
        return tttStates[0]
    }
    tttState := TTTState{
        User: user.Key,
        Started: time.Now(),
        Board: newMatrix(3),
        Key: datastore.IncompleteKey("TTTState", nil),
    }
    Client.Put(Ctx, tttState.Key, tttState)
    return &tttState
}

func getHangmanState(token string) *HangmanState {
    user := getUserBySessionToken(token)
    hangmanStates := HangmanStates{}
    Client.GetAll(Ctx, datastore.NewQuery("HangmanState").Filter("User=", user.Key), &hangmanStates)
    if len(hangmanStates) > 0 {
        return hangmanStates[0]
    }
    hangmanState := HangmanState{
        User: user.Key,
        Started: time.Now(),
        Board: newMatrix(3),
        Key: datastore.IncompleteKey("HangmanState", nil),
    }
    Client.Put(Ctx, hangmanState.Key, hangmanState)
    return &hangmanState
}

func getWWUFState(token string) *WWUFState {
    user := getUserBySessionToken(token)
    wwufStates := WWUFStates{}
    Client.GetAll(Ctx, datastore.NewQuery("WWUFState").Filter("User=", user.Key), &wwufStates)
    if len(wwufStates) > 0 {
        return wwufStates[0]
    }
    wwufState := WWUFState{
        User: user.Key,
        Started: time.Now(),
        Board: newMatrix(3),
        Key: datastore.IncompleteKey("wwufState", nil),
    }
    Client.Put(Ctx, wwufState.Key, wwufState)
    return &wwufState
}

func updateTTTState(state *TTTState) error {
    Client.Put(Ctx, state.Key, state)
    return nil
}

func updateHangmanState(state *HangmanState) error {
    Client.Put(Ctx, state.Key, state)
    return nil
}

func updateWWUFState(state *WWUFState) error {
    Client.Put(Ctx, state.Key, state)
    return nil
}

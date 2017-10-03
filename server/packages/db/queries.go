package db

import (
    "fmt"
    "time"
    "net/http"
    "cloud.google.com/go/datastore"
    "github.com/TheRoyalTnetennba/ai-wins/server/packages/utils"
)

func GetAllGames() Games {
    games := Games{}
    Client.GetAll(Ctx, datastore.NewQuery("Game"), &games)
    return games
}

// ---------------------- USER ----------------------



// ---------------------- USER ----------------------

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

func GetTTTState(user *User) *TTTState {
    tttStates := TTTStates{}
    Client.GetAll(Ctx, datastore.NewQuery("TTTState").Filter("User=", user.Key), &tttStates)
    if len(tttStates) > 0 {
        return tttStates[0]
    }
    tttState := TTTState{
        User: user.Key,
        Started: time.Now(),
        Board: utils.NewMatrix(3),
        Key: datastore.IncompleteKey("TTTState", nil),
        Game: nil,
    }
    _, err := Client.Put(Ctx, tttState.Key, tttState)
    fmt.Println(err)
    return &tttState
}

func GetHangmanState(user *User) *HangmanState {
    hangmanStates := HangmanStates{}
    Client.GetAll(Ctx, datastore.NewQuery("HangmanState").Filter("User=", user.Key), &hangmanStates)
    if len(hangmanStates) > 0 {
        return hangmanStates[0]
    }
    hangmanState := HangmanState{
        User: user.Key,
        Started: time.Now(),
        Board: utils.NewMatrix(3),
        Key: datastore.IncompleteKey("HangmanState", nil),
    }
    Client.Put(Ctx, hangmanState.Key, hangmanState)
    return &hangmanState
}

func GetWWUFState(user *User) *WWUFState {
    wwufStates := WWUFStates{}
    Client.GetAll(Ctx, datastore.NewQuery("WWUFState").Filter("User=", user.Key), &wwufStates)
    if len(wwufStates) > 0 {
        return wwufStates[0]
    }
    wwufState := WWUFState{
        User: user.Key,
        Started: time.Now(),
        Board: utils.NewMatrix(3),
        Key: datastore.IncompleteKey("wwufState", nil),
    }
    Client.Put(Ctx, wwufState.Key, wwufState)
    return &wwufState
}

func UpdateTTTState(state *TTTState) error {
    Client.Put(Ctx, state.Key, state)
    return nil
}

func UpdateHangmanState(state *HangmanState) error {
    Client.Put(Ctx, state.Key, state)
    return nil
}

func UpdateWWUFState(state *WWUFState) error {
    Client.Put(Ctx, state.Key, state)
    return nil
}

func GetToken(r *http.Request) string {
    session, _ := Store.Get(r, "ai-wins")
    return session.Values["AccessToken"].(string)
}

func GetExpiry(r *http.Request) time.Time {
    session, _ := Store.Get(r, "ai-wins")
    return session.Values["Expiry"].(time.Time)
}
package db

import (
    "fmt"
    "time"
    "net/http"
    "cloud.google.com/go/datastore"
)

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

func GetUserByOAuthID(id string, token string) Users {
    users := Users{}
    query := datastore.NewQuery("User")
    keys, _ := Client.GetAll(Ctx, query.Filter("OAuthID=", id), &users)
    if len(users) > 0 {
        users[0].Token = token
        Client.Put(Ctx, keys[0], users[0])
    }
    return users
}

func AddNewUser(user *User) error {
    newKey := datastore.IncompleteKey("User", nil)
    _, err := Client.Put(Ctx, newKey, user)
    if err != nil {
        return err
    }
    return nil
}

func deleteLocalToken(r *http.Request, w http.ResponseWriter) {
    session, _ := Store.Get(r, "ai-wins")
    session.Values["AccessToken"] = ""
    session.Values["Expiry"] = time.Now()
    session.Values["TokenType"] = ""
    session.Values["RefreshToken"] = ""
    err := session.Save(r, w)
    if err != nil {
        fmt.Println("error deleting session")
    }
}

func deleteRemoteToken(token string) {
    user := getUserBySessionToken(token)
    if len(user.Key.String()) > 0 {
        user.Token = ""
        Client.Put(Ctx, user.Key, user)
    }
}
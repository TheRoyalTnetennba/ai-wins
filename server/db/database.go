package main

import (
    "log"
    "time"
    "encoding/gob"
    "golang.org/x/net/context"
    "cloud.google.com/go/firestore"
    "github.com/gorilla/sessions"
)

func newClient() *firestore.Client {
    client, err := firestore.NewClient(ctx, ProjectID)
    if err != nil {
        log.Fatal("Failed to set up DB client", err)
    }
    return client
}

var (
    ctx context.Context = context.Background()
    client *firestore.Client = newClient()
    store = sessions.NewCookieStore([]byte(""))
    // GameKeys = make(map[string]*firestore.Key)
)

func init() {
    store.Options.MaxAge = 86400
    store.Options.Domain = "localhost"
    gob.Register(time.Time{})
    // games := GetAllGames()
    // for i := 0; i < len(games); i++ {
    //     GameKeys[games[i].Slug] = games[i].Key
    // }
}
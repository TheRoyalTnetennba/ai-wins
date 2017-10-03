package db

import (
    "log"
    "time"
    "encoding/gob"
    "golang.org/x/net/context"
    "cloud.google.com/go/datastore"
    "github.com/gorilla/sessions"
)

func newClient() *datastore.Client {
    client, err := datastore.NewClient(Ctx, ProjectID)
    if err != nil {
        log.Fatal("Failed to set up DB client", err)
    }
    return client
}

func init() {
    Store.Options.MaxAge = 86400
    Store.Options.Domain = "localhost"
    gob.Register(time.Time{})
}

var (
    ProjectID string = "personal-projects-177215"
    Ctx context.Context = context.Background()
    Client *datastore.Client = newClient()
    Store = sessions.NewCookieStore([]byte(""))
)
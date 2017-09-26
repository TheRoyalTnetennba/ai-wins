package main

import (
    "fmt"
    "log"
    "time"
    "cloud.google.com/go/datastore"
    "golang.org/x/net/context"
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
    GamesStarted int
    Lost int
    Name string
    Tied string
    Won string
}

func main() {
    ctx := context.Background()
    ProjectID := "personal-projects-177215"
    // Create a datastore client. In a typical application, you would create
    // a single client which is reused for every datastore operation.
        // Creates a client.
    client, err := datastore.NewClient(ctx, ProjectID)
    if err != nil {
            log.Fatalf("Failed to create client: %v", err)
    }

    // Sets the kind for the new entity.
    kind := "Task"
    // Sets the name/ID for the new entity.
    name := "sampletask1"
    // Creates a Key instance.
    taskKey := datastore.NameKey(kind, name, nil)

    // Creates a Task instance.
    task := Task{
            Description: "Buy milk",
    }

    // Saves the new entity.
    if _, err := client.Put(ctx, taskKey, &task); err != nil {
            log.Fatalf("Failed to save task: %v", err)
    }

    fmt.Printf("Saved %v: %v\n", taskKey, task.Description)
}

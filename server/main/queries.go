package main

import (
    "fmt"
    "cloud.google.com/go/datastore"
)

func getAllGames() Games {
    games := Games{}
    _, err1 := Client.GetAll(Ctx, datastore.NewQuery("Game"), &games)
    if err1 != nil {
      fmt.Println("Error fetching data")
    }
    for i := 0; i < len(games); i++ {
        fmt.Println(games[i])
    }
    return games
}
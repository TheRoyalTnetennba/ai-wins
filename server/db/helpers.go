package main

import (
"fmt"
// "cloud.google.com/go/firestore"
)

// ----------- ADD LIMITS TO QUERIES WHERE APPROPRIATE -------------
// ----------- ONLY ALTER REQUIRED FIELDS RATHER THAN RE-SETTING DOCUMENT -------------

func getUserByUID(uid string) User {
    user := User{}
    doc, _ := client.Doc(fmt.Sprintf("users/%s", uid)).Get(ctx)
    doc.DataTo(&user)
    return user
}

func addUser(u User) {
    client.Collection("users").Doc(u.UID).Set(ctx, &u)
}

func normalizeUser(u User) User {
    if u.TTT.Board == nil {
        u.TTT.Board = genBoard(3, 3)
    }
    return u
}

func updateUser(u User) {
    client.Collection("users").Doc(u.UID).Set(ctx, &u)
}

func setSessionToken(u User, token string) {
    u.Token = token
    doc := client.Doc(fmt.Sprintf("users/%s", u.UID))
    doc.UpdateStruct(ctx, []string{"Token"}, u)
}

func getUserBySessionToken(u User) User {
    user := User{}
    iter := client.Collection("users").Where("Token", "==", u.Token).Documents(ctx)
    for {
        doc, err := iter.Next()
        if err != nil {
            return user
        }
        doc.DataTo(&user)
    }
    return user
}

func getGameFromSlug(path string) Game {
    game := Game{}
    doc, _ := client.Doc(fmt.Sprintf("games/%s", path)).Get(ctx)
    doc.DataTo(&game)
    return game
}
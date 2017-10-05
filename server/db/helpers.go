package main

import (
// "fmt"
)

func getUserByUID(uid string) User {
    user := User{}
    iter := client.Collection("users").Where("uid", "==", uid).Documents(ctx)
    for {
        doc, err := iter.Next()
        if err != nil {
            return user
        }
        doc.DataTo(&user)
    }
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

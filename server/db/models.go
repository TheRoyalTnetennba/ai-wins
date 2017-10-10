package main

import (
    "time"
)

// ----------- DO NOT EXPORT SPECIAL FIELDS -----------

type GoogleUser struct {
    Name string `json:"displayName"`
    Email string `json:"email"`
    PhoneNumber string `json:"phoneNumber"`
    Picture string `json:"photoURL"`
    UID string `json:"uid"`
    Token OToken `json:"token"`
}

type User struct {
    UID string `json:"uid,omitempty"`
    Username string `json:"username,omitempty"`
    Image string `json:"image,omitempty"`
    Email string `json:"email,omitempty"`
    PhoneNumber string `json:"phoneNumber,omitempty"`
    Joined time.Time `json:"joined,omitempty"`
    Lost int `json:"lost,omitempty"`
    Tied int `json:"tied,omitempty"`
    Won int `json:"won,omitempty"`
    Wallet string `json:"wallet,omitempty"`
    Mined string `json:"mined,omitempty"`
    Token OToken `json:"token",omitempty`
    TTT TTTState `json:"ttt,omitempty"`
    WWUF WWUFState `json:"wwuf,omitempty"`
    Hangman HangmanState `json:"hangman,omitempty"`
}

type OToken struct {
    AccessToken string `json:"accessToken,omitempty"`
    IDToken string `json:"idToken,omitempty"`
    ProviderID string `json:"providerID,omitempty"`
}

type Game struct {
    Description string `json:"description", firestore:"description"`
    Image string `json:"image", firestore:"image"`
    DateAdded time.Time `json:"dateAdded", firestore:"date_added"`
    LastLoss time.Time `json:"lastLoss", firestore:"last_loss"`
    LastPayout string `json:"lastPayout", firestore:"last_payout"`
    CanTie bool `json:"canTie", firestore:"can_tie"`
    LastWinner string `json:"lastWinner", firestore:"last_winner"`
    Tagline string `json:"tagline", firestore:"tagline"`
    Name string `json:"name", firestore:"name"`
    Lost int `json:"lost", firestore:"lost"`
    Tied int `json:"tied", firestore:"tied"`
    Won int `json:"won", firestore:"won"`
    Slug string `json:"slug", firestore:"slug"`
    Consecutive int `json:"consecutive", firestore:"consecutive"`
}

type WWUFState struct {
    Started time.Time `json:"started", firestore:"started"`
    Board map[string][]string `json:"board", firestore:"board"`
    UserLetters []string `json:"userLetters", firestore:"user_letters"`
    AILetters []string `json:"aiLetters", firestore:"ai_letters"`
    B1 string `json:"b1", firestore:"b1"`
    B2 string `json:"b2", firestore:"b2"`
    RemainingLetters []string `json:"remaingLetters", firestore:"remaing_letters"`
}

type Games []Game

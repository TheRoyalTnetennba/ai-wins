package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GoogleLogin",
		"GET",
		"/api/v1/googlelogin",
		GoogleLogin,
	},
	Route{
		"GoogleCallback",
		"GET",
		"/api/v1/googlecallback",
		GoogleCallback,
	},
	Route{
		"GameData",
		"GET",
		"/api/v1/games/{gameID}",
		GameData,
	},
	Route{
		"GetMove",
		"POST",
		"/api/v1/games/{gameName}/getMove",
		GetMove,
	},
	Route{
		"GamesIndex",
		"GET",
		"/api/v1/games",
		GamesIndex,
	},
	Route{
		"WWUFGetLetters",
		"GET",
		"/api/v1/wwuf/getletters/{numLetters}",
		WWUFGetLetters,
	},
}

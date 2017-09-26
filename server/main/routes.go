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
		"Register",
		"GET",
		"/register",
		Register,
	},
	Route{
		"GoogleLogin",
		"GET",
		"/googlelogin",
		GoogleLogin,
	},
	Route{
		"GoogleCallback",
		"GET",
		"/googlecallback",
		GoogleCallback,
	},
	Route{
		"GameShow",
		"GET",
		"/games/{gameID}",
		GameShow,
	},
	Route{
		"GetMove",
		"POST",
		"/games/getMove",
		GetMove,
	},
	Route{
		"GamesIndex",
		"GET",
		"/games",
		GamesIndex,
	},
}

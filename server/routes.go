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
		"GetUser",
		"GET",
		"/api/v1/sec/{resource}",
		GetUser,
	},
	Route{
		"tttState",
		"POST",
		"/api/v1/sec/ttt",
		tttState,
	},
	Route{
		"GameData",
		"GET",
		"/api/v1/{resource}",
		GameData,
	},
	// Route{
	// 	"WWUFGetLetters",
	// 	"GET",
	// 	"/api/v1/wwuf/getletters/{numLetters}",
	// 	WWUFGetLetters,
	// },
}

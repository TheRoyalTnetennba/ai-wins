package main

import (
	"fmt"
	"net/http"
	"encoding/json"
    simplejson "github.com/bitly/go-simplejson"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
    "time"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/googlecallback", BaseURL),
		ClientID:     Google.Key,
		ClientSecret: Google.Secret,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	oauthStateString = randSeq(13)
)

func updateSession(r *http.Request, w http.ResponseWriter, token *oauth2.Token) {
    session, _ := Store.Get(r, "ai-wins")
    session.Values["sessionToken"] = token
    session.Save(r, w)
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	res := simplejson.New()
	res.Set("url", url)
	payload, _ := res.Encode()
	w.Write(payload)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
    updateSession(r, w, token)

    response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
    defer response.Body.Close()
    var gUser GoogleUser
    json.NewDecoder(response.Body).Decode(&gUser)
    users := getUserByOAuthID(gUser.ID)
    if len(users) < 1 {
        newUser := User{
            Image: gUser.Picture,
            Joined: time.Now(),
            Lost: 0,
            Tied: 0,
            UserName: gUser.Name,
            Won: 0,
            Email: gUser.Email,
            Token: token.AccessToken,
            OAuthID: gUser.ID,
        }
        go addNewUser(&newUser)
    }
    http.Redirect(w, r, fmt.Sprintf("%s/#/auth-callback", ClientURL), http.StatusTemporaryRedirect)
}

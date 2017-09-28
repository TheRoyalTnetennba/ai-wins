package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	simplejson "github.com/bitly/go-simplejson"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/TheRoyalTnetennba/ai-wins/server/utils"
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
	oauthStateString = utils.RandSeq(13)
)

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
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Content: %s\n", contents)
}

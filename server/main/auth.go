package main

import (
	"fmt"
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
    session, err := Store.Get(r, "ai-wins")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println(token)
    res, _ := simplejson.NewFromReader(response.Body)
// {"email":"gpaye8@gmail.com","family_name":"Paye","gender":"male","given_name":"Graham","id":"112193450734641047471","link":"https://plus.google.com/112193450734641047471","locale":"en","name":"Graham Paye","picture":"https://lh5.googleusercontent.com/-LbON--7wLOI/AAAAAAAAAAI/AAAAAAAAFnk/0xR3UM6iQYQ/photo.jpg","verified_email":true}
    payload, _ := res.Encode()
    session.Values["babanas"] = "bananas"
    session.Save(r, w)
	fmt.Fprintf(w, "Content: %s\n", payload)
}

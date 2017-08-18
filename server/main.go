package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "io/ioutil"
)
// "github.com/gorilla/mux" add to import
//  "os"

const htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
</body></html>
`

func handleMain(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, htmlIndex)
}

filePath := "./key.json"

var (
    googleOauthConfig = &oauth2.Config{
        RedirectURL:    "http://localhost:3000/GoogleCallback",
        // ClientID:     os.Getenv("googlekey"),
        ClientID:     "687671941121-lnojr5i0f2eknrdpivcdah6pfhiruro7.apps.googleusercontent.com",
        // ClientSecret: os.Getenv("googlesecret"),
        ClientSecret: "bHG1-oSVkf0MOs9QSwTJvS_P",
        Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile",
            "https://www.googleapis.com/auth/userinfo.email"},
        Endpoint:     google.Endpoint,
    }
// Some random string, random for each request
    oauthStateString = "random"
)


func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
    url := googleOauthConfig.AuthCodeURL(oauthStateString)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
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

func main() {
    http.HandleFunc("/", handleMain)
    http.HandleFunc("/GoogleLogin", handleGoogleLogin)
    http.HandleFunc("/GoogleCallback", handleGoogleCallback)
    fmt.Println(http.ListenAndServe(":3000", nil))
}

package wwuf

func WWUFGetLetters(w http.ResponseWriter, r *http.Request) {
    req, _ := simplejson.NewFromReader(r.Body)
    ch := make(chan []byte)
    go GetAIMove(req.MustMap(), ch)
    payload := <-ch
    w.Header().Set("Content-Type", "application/json")
    w.Write(payload)
}
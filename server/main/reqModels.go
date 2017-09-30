type GoogleUser struct {
    Email string
    Family_name string
    Gender string
    Given_name string
    Id string
    Link string
    Locale string
    Name string
    Picture string
    Verified_email bool
}

type TTTMove struct {
    Board [][]string
    Marker string
    Move [] int
}

type HangmanMove struct {
    Board [][]string

}

func GetAIMove(body map[string]interface{}, c chan []byte) {
    board := utils.GetMatrixFromInterface(body["gameState"])
    pos := GetBestMove(board, body["marker"].(string))
    res := simplejson.New()
    res.Set("move", pos)
    payload, _ := res.Encode()
    c <- payload
}

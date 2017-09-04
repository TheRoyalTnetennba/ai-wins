package main

import (
	"github.com/TheRoyalTnetennba/ai-wins/server/games/tic-tac-toe"
	"github.com/TheRoyalTnetennba/ai-wins/server/utils"
	simplejson "github.com/bitly/go-simplejson"
)

// "crypto/sha512"
// ticTacToe := fmt.Sprintf("%x", sha512.Sum512([]byte("ticTacToe")))

func GetAIMove(body map[string]interface{}, c chan []byte) {
	// board := utils.GetMatrixFromInterface(body["gameState"])
	board := utils.GetMatrixFromInterface(body["gameState"])
	pos := tttAI.GetAIMove(board)
	res := simplejson.New()
	res.Set("move", pos)
	// gameID := (req.Get("gameID")).MustString()
	// marker := (req.Get("marker")).MustArray()
	payload, _ := res.Encode()
	c <- payload
}

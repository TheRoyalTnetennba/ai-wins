package main

import (
	"github.com/TheRoyalTnetennba/ai-wins/server/ai"
	"github.com/TheRoyalTnetennba/ai-wins/server/utils"
	simplejson "github.com/bitly/go-simplejson"
)

// "crypto/sha512"
// ticTacToe := fmt.Sprintf("%x", sha512.Sum512([]byte("ticTacToe")))

func GetAIMove(body map[string]interface{}, c chan []byte) {
	board := utils.GetMatrixFromInterface(body["gameState"])
	pos := tttAI.GetAIMove(board, body["marker"].(string))
	res := simplejson.New()
	res.Set("move", pos)
	payload, _ := res.Encode()
	c <- payload
}

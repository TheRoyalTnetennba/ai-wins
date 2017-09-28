package utils

import (
	"math/rand"
	"time"
)

var (
	Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	rand.Seed(time.Now().Unix())
}

// GetMatrixFromInterface processes simplejson-created interfaces of a stringified matrix
func GetMatrixFromInterface(inter interface{}) [][]string {
	var matrix [][]string
	arr := inter.([]interface{})
	for i := 0; i < len(arr); i++ {
		var row []string
		arrRow := arr[i].([]interface{})
		for j := 0; j < len(arrRow); j++ {
			row = append(row, arrRow[j].(string))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func CopyMatrix(orig [][]string) [][]string {
	var newMatrix [][]string
	for i := 0; i < len(orig); i++ {
		var row []string
		for j := 0; j < len(orig[i]); j++ {
			row = append(row, orig[i][j])
		}
		newMatrix = append(newMatrix, row)
	}
	return newMatrix
}

func RandNum(upperLimit int) int {
	return randIntn(upperLimit)
}

func RandSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = Letters[rand.Intn(len(Letters))]
    }
    return string(b)
}
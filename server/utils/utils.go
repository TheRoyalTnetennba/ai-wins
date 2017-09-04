package utils

import (
	"fmt"
	"reflect"
)

// GetMatrixFromInterface processes simplejson-created interfaces of a stringified matrix
func GetMatrixFromInterface(inter interface{}) [][]string {
	var matrix [][]string
	for i := 0; i < 3; i++ {
		var row []string
		for j := 0; j < 3; j++ {
			row = append(row, "")
		}
		matrix = append(matrix, row)
	}
	// length := reflect.ValueOf(inter).Len()
	arr := reflect.Value(reflect.ValueOf(inter))
	fmt.Println(arr.Index(0))
	return arr
}

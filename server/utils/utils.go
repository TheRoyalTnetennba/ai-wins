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
	arr := reflect.ValueOf(inter)
	for i := 0; i < arr.Len(); i++ {
		for j := 0; j < arr.Len(); j++ {
			fmt.Println(reflect.ValueOf(arr.Index(i)))
		}
	}
	return matrix
}

package utils

import (
	"reflect"
)

// GetMatrixFromInterface processes simplejson-created interfaces of a stringified matrix
func GetMatrixFromInterface(inter interface{}) [][]string {
	var matrix [][]string
	arr := reflect.ValueOf(inter)
	for i := 0; i < arr.Len(); i++ {
		var row []string
		col := reflect.ValueOf(arr.Index(i))
		for j := 0; j < col.Len(); j++ {
			row = append(row, arr.Index(i).Index(j).String())
		}
		matrix = append(matrix, row)
	}
	return matrix
}

package utils

import (
	"reflect"
)

func GetMatrixFromInterface(inter interface{}) [][]string {
	var matrix [][]string
	arr := reflect.ValueOf(inter)
	for i := 0; i < arr.Len(); i++ {
		var row []string
		for j := 0; j < arr.Index(i).Len(); j++ {
			row = append(row, arr.Index(i).Index(j).String())
		}
		matrix = append(matrix, row)
	}
	return matrix
}

package utils

import "fmt"

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
	// arr := reflect.Value(reflect.ValueOf(inter))
	// for i := 0; i < arr.Len(); i++ {
	// 	rowInt := reflect.Value(reflect.ValueOf(arr))
	// 	fmt.Println(rowInt)
	// 	// var row []string
	// 	// for j := 0; j < rowInt.Len(); j++ {
	// 	// 	fmt.Println(rowInt.Index(j), rowInt.Kind())
	// 	// }
	// }
	arr := inter.([]interface{})
	row := arr[0].([]interface{})
	square := row[2].(string)
	fmt.Println(square)
	return matrix
}

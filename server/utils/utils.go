package utils

import "fmt"

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
	// arr := reflect.Value(reflect.ValueOf(inter))
	// for i := 0; i < arr.Len(); i++ {
	// 	rowInt := reflect.Value(reflect.ValueOf(arr))
	// 	fmt.Println(rowInt)
	// 	// var row []string
	// 	// for j := 0; j < rowInt.Len(); j++ {
	// 	// 	fmt.Println(rowInt.Index(j), rowInt.Kind())
	// 	// }
	// }
	fmt.Println(matrix)
	return matrix
}

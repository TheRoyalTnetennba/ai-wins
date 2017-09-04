package utils

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

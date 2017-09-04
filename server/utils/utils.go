package utils

func getMatrixFromInterface(arr []interface{}) [][]string {
	var matrix [][]string
	for _, i := range arr {
		matrix = append(matrix, i.([]string))
	}
	return matrix
}

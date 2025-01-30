package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

func main() {
	mtxs := createMatrix(10, 10)
	for _, mtx := range mtxs {
		fmt.Println(mtx)
	}
}

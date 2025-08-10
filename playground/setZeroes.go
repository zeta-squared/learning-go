package main

func setZeroes(matrix [][]int) {
	zeroIndices := [][]int{}
	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			if matrix[i][j] == 0 {
				zeroIndices = append(zeroIndices, []int{i, j})
			}
		}
	}

	for _, v := range zeroIndices {
		// Iterate over the entire row to set zeroes.
		for j := range len(matrix[0]) {
			matrix[v[0]][j] = 0
		}

		// Iterate over the entire column to set zeroes.
		for i := range len(matrix) {
			matrix[i][v[1]] = 0
		}
	}
}

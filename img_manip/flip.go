// Package img_manip implements functions for manipulating images.
package img_manip

import "fmt"

// FlipUD flips the provided matrix vertically.
func FlipUD(matrix [][]int) ([][]int, error) {
	num_rows, num_cols := len(matrix), len(matrix[0])

	if 28 < num_rows || num_rows < 1 {
		return nil, fmt.Errorf("number of rows must be between 1 and 28")
	}
	if 28 < num_cols || num_cols < 1 {
		return nil, fmt.Errorf("number of columns must be between 1 and 28")
	}

	flipped_matrix := make([][]int, num_rows)
	for row := range flipped_matrix {
		flipped_matrix[row] = make([]int, num_cols)
	}

	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			flipped_matrix[row][col] = matrix[num_rows-row-1][col]
		}
	}

	return flipped_matrix, nil
}

// FlipLR flips the provided matrix horizontally.
func FlipLR(matrix [][]int) ([][]int, error) {
	num_rows, num_cols := len(matrix), len(matrix[0])

	if 28 < num_rows || num_rows < 1 {
		return nil, fmt.Errorf("number of rows must be between 1 and 28")
	}
	if 28 < num_cols || num_cols < 1 {
		return nil, fmt.Errorf("number of columns must be between 1 and 28")
	}

	flipped := make([][]int, num_rows)
	for row := range flipped {
		flipped[row] = make([]int, num_cols)
	}

	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			flipped[row][col] = matrix[row][num_cols-col-1]
		}
	}

	return flipped, nil
}

// FlipAllAxes flips the provided matrix horizontally and vertically.
func FlipAllAxes(matrix [][]int) ([][]int, error) {
	num_rows, num_cols := len(matrix), len(matrix[0])

	if 28 < num_rows || num_rows < 1 {
		return nil, fmt.Errorf("number of rows must be between 1 and 28")
	}
	if 28 < num_cols || num_cols < 1 {
		return nil, fmt.Errorf("number of columns must be between 1 and 28")
	}

	flipped_matrix := make([][]int, num_rows)
	for row := 0; row < num_rows; row++ {
		flipped_matrix[row] = make([]int, num_cols)
	}

	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			flipped_matrix[row][col] = matrix[num_rows-row-1][num_cols-col-1]
		}
	}

	return flipped_matrix, nil
}

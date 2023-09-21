// Package img_manip implements functions for manipulating images.
package img_manip

// Rotate90 rotates the provided matrix 90 degrees clockwise the provided
// number of times.
func Rotate90(matrix [][]int, num_rotations int) [][]int {
	num_rotations = num_rotations % 4

	if num_rotations == 0 {
		return matrix
	}

	num_rows, num_cols := len(matrix), len(matrix[0])

	rotated_matrix := make([][]int, num_cols)
	for col := range rotated_matrix {
		rotated_matrix[col] = make([]int, num_rows)
	}

	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			switch num_rotations {
			case 1:
				rotated_matrix[col][num_rows-row-1] = matrix[row][col]
			case 2:
				rotated_matrix[num_cols-col-1][num_rows-row-1] = matrix[row][col]
			case 3:
				rotated_matrix[num_cols-col-1][row] = matrix[row][col]
			}
		}
	}

	return rotated_matrix
}

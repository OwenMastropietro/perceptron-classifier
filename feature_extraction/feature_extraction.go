// Package feature_extraction implements functions for extracting feature
// values from images.
package feature_extraction

import (
	"fmt"
	"os"

	"project04_perceptron/go_rewrite/img_manip"
)

/*
Feature 1

GetDensity returns the density of the provided image.
*/
func GetDensity(image [][]int) float64 {
	num_rows, num_cols := len(image), len(image[0])

	sum := 0
	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			sum += image[row][col]
		}
	}

	return float64(sum) / float64(num_rows*num_cols)
}

/*
Feature 2

GetVerticalSymmetry returns the degree of vertical symmetry of the provided
image. The degree of symmetry is defined as the average greyscale value
of the image obtained by the bitwise XOR of each pixel with its
corresponding vertically reflected image.
For example, let I be the image and let I' be the image whose j-th column
is the (num_columns - j)-th column of I. Then the measure of symmetry is
the density of I XOR I'.
*/
func GetVerticalSymmetry(image [][]int) float64 {
	num_rows, num_cols := len(image), len(image[0])

	reflected_image := make([][]int, num_rows)
	for row := range reflected_image {
		reflected_image[row] = make([]int, num_cols)
	}

	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			reflected_image[row][col] = image[row][num_cols-col-1]
		}
	}

	xor_image := make([][]int, num_rows)
	for row := range xor_image {
		xor_image[row] = make([]int, num_cols)
	}

	// bitwise XOR
	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			xor_image[row][col] = image[row][col] ^ reflected_image[row][col]
		}
	}

	return GetDensity(xor_image)
}

/*
Feature 3

GetVerticalIntersections returns the maximum number of vertical
intersections and the average number of vertical intersections of the
provided image.
*/
func GetVerticalIntersections(image [][]int) (int, float64) {
	num_rows, num_cols := len(image), len(image[0])

	counts := make([]int, num_rows)
	for row := 0; row < num_rows; row++ {
		count := 0
		prev := 0
		for col := 0; col < num_cols; col++ {
			current := image[col][row]
			if current != prev {
				count++
			}
			prev = current
		}
		counts[row] = count
	}

	maximum_intersections, average_intersections := 0, 0.0
	for _, count := range counts {
		if count > maximum_intersections {
			maximum_intersections = count
		}
		average_intersections += float64(count)
	}
	average_intersections /= float64(num_rows)

	return maximum_intersections, average_intersections
}

/*
Feature 4

GetHorizontalIntersections returns the maximum number of horizontal
intersections and the average number of horizontal intersections of the
provided image.
*/
func GetHorizontalIntersections(image [][]int) (int, float64) {

	return GetVerticalIntersections(img_manip.Rotate90(image, 1))
}

/*
Feature 5

GetNumLoops returns the number of loops in the provided image by performing
a breadth-first search on the image. The black pixels are considered unvisited
and the white pixels are considered visited. The number of loops is the number
of black sections that must be flood filled to turn the entire image white.
*/
func GetNumLoops(image [][]int) (int, error) {
	num_rows, num_cols := len(image), len(image[0])

	black, white := 0, 1

	is_valid := func(x, y int) bool {
		return 0 <= x && x < num_rows && 0 <= y && y < num_cols
	}

	bfs := func(x, y int) {
		queue := [][2]int{{x, y}}

		directions := [][2]int{
			{1, -1}, {1, 0}, {1, 1},
			{0, -1} /*{.}*/, {0, 1},
			{-1, -1}, {-1, 0}, {-1, 1}}

		for len(queue) > 0 {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]
			image[x][y] = white

			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]

				if is_valid(nx, ny) && image[nx][ny] == black {
					queue = append(queue, [2]int{nx, ny})
					image[nx][ny] = white
				}

			}
		}

	}

	search_count := 0
	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			if image[row][col] == black {
				bfs(row, col)
				search_count++
			}
		}
	}

	return search_count - 1, nil
}

/*
Feature 6

GetUDSymmetry returns the degree of symmetry of the provided image
by reflecting the top half of the image over the bottom half and
computing the density of the bitwise XOR of the two images.
*/
func GetUDSymmetry(image [][]int) float64 {
	num_rows, num_cols := len(image), len(image[0])
	num_rows_split := num_rows / 2

	top := image[:num_rows_split]
	bottom, err := img_manip.FlipUD(image[num_rows_split:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	xor_image := make([][]int, num_rows_split)
	for row := range xor_image {
		xor_image[row] = make([]int, num_cols)
	}

	for row := 0; row < num_rows_split; row++ {
		for col := 0; col < num_cols; col++ {
			xor_image[row][col] = top[row][col] ^ bottom[row][col]
		}
	}

	return GetDensity(xor_image)
}

/*
Feature 7

GetLRSymmetry returns the degree of symmetry of the provided image
by rotating the image 90 degrees clockwise and calling GetUDSymmetry
as if to reflect the left half of the image over the right half and
compute the density of the bitwise XOR of the two images.
*/
func GetLRSymmetry(image [][]int) float64 {

	return GetUDSymmetry(img_manip.Rotate90(image, 1))
}

/*
Feature 8

GetFeatureValues returns the feature values of the provided image.
*/
func GetFeatureValues(image [][]int) []float64 {
	num_features := 9
	feature_values := make([]float64, num_features)

	// GetDensity
	feature_values[0] = GetDensity(image)

	// GetVerticalSymmetry
	feature_values[1] = GetVerticalSymmetry(image)

	// GetVerticalIntersections
	max_vert, avg_vert := GetVerticalIntersections(image)
	feature_values[2] = float64(max_vert)
	feature_values[3] = avg_vert

	// GetHorizontalIntersections
	max_horiz, avg_horiz := GetHorizontalIntersections(image)
	feature_values[4] = float64(max_horiz)
	feature_values[5] = avg_horiz

	// GetNumLoops
	num_loops, err := GetNumLoops(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	feature_values[6] = float64(num_loops)

	// GetUDSymmetry
	feature_values[7] = GetUDSymmetry(image)

	// GetLRSymmetry
	feature_values[8] = GetLRSymmetry(image)

	return feature_values
}

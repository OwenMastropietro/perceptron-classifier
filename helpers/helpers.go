// Package helpers implements helper functions for the perceptron model.
package helpers

import (
	"bufio"
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// DotProduct returns the dot product of the provided vectors.
func DotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("Dot product of vectors of unequal length")
	}

	result := 0.0
	for i := range a {
		result += a[i] * b[i]
	}

	return result
}

// ArgMax returns the index of the maximum value in the provided slice.
func ArgMax(a []float64) int {
	max_value, max_index := a[0], 0

	for i := range a {
		if a[i] > max_value {
			max_value, max_index = a[i], i
		}
	}

	return max_index
}

// ArgMin returns the index of the minimum value in the provided slice.
func ArgMin(a []int) int {
	min_value, min_index := a[0], 0

	for i := range a {
		if a[i] < min_value {
			min_value, min_index = a[i], i
		}
	}

	return min_index
}

// AddVectors returns the element-wise sum of the provided vectors.
func AddVectors(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("Addition of vectors of unequal length")
	}

	result := make([]float64, len(a))
	for i := range a {
		result[i] = a[i] + b[i]
	}

	return result
}

// SubtractVectors returns the element-wise difference of the provided vectors.
func SubtractVectors(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("Subtraction of vectors of unequal length")
	}

	result := make([]float64, len(a))
	for i := range a {
		result[i] = a[i] - b[i]
	}

	return result
}

// Multiply returns the element-wise product of the provided vector and scalar.
func Multiply(vector []float64, scalar float64) []float64 {
	result := make([]float64, len(vector))
	for i := range vector {
		result[i] = vector[i] * scalar
	}
	return result
}

// ExtractImages reads the provided CSV file and returns a slice of images
// and a slice of labels if has_label is true. If has_label is false, the
// returned slice of labels will be nil.
func ExtractImages(file string, has_label bool) ([][][]int, []int, error) {
	// Open the CSV file
	csv_file, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer csv_file.Close()

	// Create a CSV reader
	reader := csv.NewReader(csv_file)

	first_line, err := reader.Read()
	if err != nil {
		return nil, nil, err
	}

	first_line[0] = strings.TrimPrefix(first_line[0], "\ufeff")

	// Read all CSV records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	// Initialize slices for images and labels
	var images [][][]int
	var labels []int

	// Determine the starting column index based on whether the CSV has labels
	start_col := 0
	if has_label {
		start_col = 1
	}

	// Iterate over each row in the CSV
	for _, record := range records {
		// Parse the label if it exists
		if has_label {
			label, err := strconv.Atoi(record[0])
			if err != nil {
				return nil, nil, err
			}
			labels = append(labels, label)
		}

		// Parse the pixel values and convert to a 28x28 image
		image := make([][]int, 28)
		for i := 0; i < 28; i++ {
			image[i] = make([]int, 28)
			for j := 0; j < 28; j++ {
				pixelValue, err := strconv.Atoi(record[start_col+i*28+j])
				if err != nil {
					return nil, nil, err
				}
				image[i][j] = pixelValue
			}
		}
		images = append(images, image)
	}

	return images, labels, nil
}

// GetBlackWhite returns the image after forcing all pixels less than the
// threshold to black and all pixels greater than or equal to the threshold
// to white.
func GetBlackWhite(image [][]int, threshold int) [][]int {
	black, white := 0, 1

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] < threshold {
				image[i][j] = black
			} else {
				image[i][j] = white
			}
		}
	}
	return image
}

// GetRandomWeights returns a slice of random weights with the provided
// shape and between -0.05 and 0.05.
func GetRandomWeights(rows, cols int) [][]float64 {
	weights := make([][]float64, rows)
	for i := range weights {
		weights[i] = make([]float64, cols)
		for j := range weights[i] {
			weights[i][j] = rand.Float64()*0.1 - 0.05
		}
	}
	return weights
}

// ReadIntegersFromFile reads the provided file, line byt line, and returns a
// slice of integers.
func ReadIntegersFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var integers []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return integers, nil
}

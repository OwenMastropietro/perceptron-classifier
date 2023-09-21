// Package testing_framework implements functions for testing the feature
// extraction functions.
package testing_framework

import (
	"fmt"
	"testing"

	"project04_perceptron/go_rewrite/feature_extraction"
	"project04_perceptron/go_rewrite/helpers"
)

// TestNumLoops tests the GetNumLoops feature extraction function by comparing
// the results of the function to the expected results from a file.
func TestNumLoops(t *testing.T) {
	filename := "../input_files/expected_output/num_loops.txt"
	expected, err := helpers.ReadIntegersFromFile(filename)
	if err != nil {
		t.Error(err)
	}

	actual := []int{}
	for i := 0; i < 10; i++ {
		filename := fmt.Sprintf("../input_files/training_data/handwritten_samples_%d.csv", i)

		images, _, err := helpers.ExtractImages(filename, true)
		if err != nil {
			t.Error(err)
		}

		for i := range images {
			binary_image := helpers.GetBlackWhite(images[i], 128)
			num_loops, err := feature_extraction.GetNumLoops(binary_image)
			if err != nil {
				t.Error(err)
			}
			actual = append(actual, num_loops)
		}
	}

	if len(expected) != len(actual) {
		t.Errorf("expected %d elements, got %d elements", len(expected), len(actual))
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("image %d: expected %d, got %d", i, expected[i], actual[i])
		}
	}
}

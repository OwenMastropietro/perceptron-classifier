// Package display implements functions for displaying images, statistics,
// and feature values in a "pretty" format.
package display

import "fmt"

// PrintAllImages prints the label followed by a call to PrintImage for each
// of the labels and images in the provided slices.
func PrintAllImages(labels []int, images [][][]int) {
	for i := 0; i < len(images); i++ {
		fmt.Println("Label:", labels[i])
		PrintImage(images[i])
	}
}

// PrintImage prints the provided image using whitespace X's to represent black
// and white pixels, respectively.
func PrintImage(image [][]int) {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			// fmt.Printf("%3d", image[i][j])
			if image[i][j] == 0 {
				fmt.Print("-")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}

// Stats will print the best weights, the total number of successful and
// unsuccessful predictions, and the relative success rate of said predictions.
func Stats(weights [][]float64, num_successes, num_errors, epochs int) {
	success_rate := 100 * (float64(num_successes) / float64(num_successes+num_errors))
	error_rate := 100 * (float64(num_errors) / float64(num_errors+num_successes))

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Best Weights:")
	fmt.Println(weights)

	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Relative Success (over %d Epochs on the validation data):\n", epochs)
	fmt.Printf("- Success Rate: %f%%. ", success_rate)
	fmt.Printf("(from %d successful predictions)\n", num_successes)
	fmt.Printf("- Error Rate: %f%%. ", error_rate)
	fmt.Printf("(from %d unsuccessful predictions)\n", num_errors)
}

// PPFeatureValues "pretty prints" the provided feature values, printing the
// name of the feature followed by its value.
func PPFeatureValues(feature_values []float64) {
	fmt.Printf("Density:\t\t\t%f\n", feature_values[0])
	fmt.Printf("Vertical Symmetry:\t\t%f\n", feature_values[1])
	fmt.Printf("Max Vertical Intersections:\t%f\n", feature_values[2])
	fmt.Printf("Avg Vertical Intersections:\t%f\n", feature_values[3])
	fmt.Printf("Max Horizontal Intersections:\t%f\n", feature_values[4])
	fmt.Printf("Avg Horizontal Intersections:\t%f\n", feature_values[5])
	fmt.Printf("Number of Loops:\t\t%f\n", feature_values[6])
	fmt.Printf("Vertically Split Symmetry:\t%f\n", feature_values[7])
	fmt.Printf("Horizontally Split Symmetry:\t%f\n", feature_values[8])
	fmt.Println()
}

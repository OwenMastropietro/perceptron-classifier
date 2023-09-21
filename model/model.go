// Package model implements functions for training and testing a perceptron
// model.
package model

import (
	"fmt"
	"math/rand"

	"project04_perceptron/go_rewrite/feature_extraction"
	"project04_perceptron/go_rewrite/helpers"
)

// GetTrainingData returns a slice of slices of floats representing the training
// data. Each row of this training data structure represents a single image.
// The first len(row)-2 columns represent the feature values for the image.
// The second to last column represents the threshold value (-1) for the image.
// The last column represents the class label for the image.
func GetTrainingData(verbose bool) ([][]float64, error) {
	if verbose {
		fmt.Println("------------------------------------------------")
		fmt.Println("Building Training Set...")
	}

	training_data := [][]float64{}
	class_labels := [][]int{}

	num_files := 10
	for i := 0; i < num_files; i++ {
		filename := fmt.Sprintf("input_files/training_data/handwritten_samples_%d.csv", i)
		images, labels, err := helpers.ExtractImages(filename, true)
		if err != nil {
			return nil, err
		}

		for _, label := range labels {
			class_labels = append(class_labels, []int{label})
		}

		if verbose {
			fmt.Printf("\tComputing Feature Values in < %s >...\n", filename)
		}

		for _, image := range images {
			binary_image := helpers.GetBlackWhite(image, 128)
			feature_values := feature_extraction.GetFeatureValues(binary_image)
			training_data = append(training_data, feature_values)
		}
	}

	// Create column 11 of threshold values = -1.
	// 9,990 instead of 10,000 because we dropped the 10 labels.
	thresholds := make([][]float64, len(training_data))
	for i := range thresholds {
		thresholds[i] = []float64{-1}
	}

	// Concatenate the threshold values to the training data.
	for i := range training_data {
		training_data[i] = append(training_data[i], thresholds[i]...)
	}

	// Concatenate the class labels to the training data.
	for i := range training_data {
		training_data[i] = append(training_data[i], float64(class_labels[i][0]))
	}

	rand.Shuffle(len(training_data), func(i, j int) {
		training_data[i], training_data[j] = training_data[j], training_data[i]
	})

	if len(training_data) != 9990 {
		return nil, fmt.Errorf("training data length is %d, expected 9990", len(training_data))
	}
	if len(training_data[0]) != 11 {
		return nil, fmt.Errorf("training data width is %d, expected 11", len(training_data[0]))
	}

	return training_data, nil
}

// GetValidationData returns a slice of slices of floats representing the
// validation data. Each row of this validation data structure represents a
// single image.
// The first len(row)-2 columns represent the feature values for the image.
// The second to last column represents the threshold value (-1) for the image.
// The last column represents the class label for the image.
func GetValidationData(verbose bool) ([][]float64, error) {
	if verbose {
		fmt.Println("------------------------------------------------")
		fmt.Println("Building Validation Set...")
	}

	validation_data := [][]float64{}
	class_labels := [][]int{}

	num_files := 10
	for i := 0; i < num_files; i++ {
		filename := fmt.Sprintf("input_files/validation_data/handwritten_samples_%d.csv", i)
		images, labels, err := helpers.ExtractImages(filename, true)
		if err != nil {
			return nil, err
		}

		for _, label := range labels {
			class_labels = append(class_labels, []int{label})
		}

		if verbose {
			fmt.Printf("\tComputing Feature Values in < %s >...\n", filename)
		}

		for _, image := range images {
			binary_image := helpers.GetBlackWhite(image, 128)
			feature_values := feature_extraction.GetFeatureValues(binary_image)
			validation_data = append(validation_data, feature_values)
		}
	}

	// Create column 11 of threshold values = -1.
	// 9,990 instead of 10,000 because we dropped the 10 labels.
	thresholds := make([][]float64, len(validation_data))
	for row := range thresholds {
		thresholds[row] = []float64{-1}
	}

	// Concatenate the threshold values to the training data.
	for row := range validation_data {
		validation_data[row] = append(validation_data[row], thresholds[row]...)
	}

	// Concatenate the class labels to the training data.
	for row := range validation_data {
		validation_data[row] = append(validation_data[row], float64(class_labels[row][0]))
	}

	rand.Shuffle(len(validation_data), func(i, j int) {
		validation_data[i], validation_data[j] = validation_data[j], validation_data[i]
	})

	if len(validation_data) != 2490 {
		return nil, fmt.Errorf("validation data length is %d, expected 2490", len(validation_data))
	}
	if len(validation_data[0]) != 11 {
		return nil, fmt.Errorf("validation data width is %d, expected 11", len(validation_data[0]))
	}

	return validation_data, nil
}

// GetTestingData returns a slice of slices of floats representing the testing
// data. Each row of this testing data structure represents a single image
// extracted from the provided file.
// The first len(row)-1 columns represent the feature values for the image.
// The last column represents the threshold value (-1) for the image.
// The class label for each image is omitted.
func GetTestingData(file string, verbose bool) ([][]float64, error) {
	if verbose {
		fmt.Println("------------------------------------------------")
		fmt.Println("Building Testing Set...")
		fmt.Printf("\tComputing Feature Values in < %s >...\n", file)
	}

	testing_data := [][]float64{}

	images, _, err := helpers.ExtractImages(file, false)
	if err != nil {
		return nil, err
	}

	for _, image := range images {
		binary_image := helpers.GetBlackWhite(image, 128)
		feature_values := feature_extraction.GetFeatureValues(binary_image)
		testing_data = append(testing_data, feature_values)
	}

	thresholds := make([][]float64, len(testing_data))
	for row := range thresholds {
		thresholds[row] = []float64{-1}
	}

	for row := range testing_data {
		testing_data[row] = append(testing_data[row], thresholds[row]...)
	}

	if len(testing_data) != 99 {
		return nil, fmt.Errorf("testing data length is %d, expected 99", len(testing_data))
	}
	if len(testing_data[0]) != 10 {
		return nil, fmt.Errorf("testing data width is %d, expected 11", len(testing_data[0]))
	}

	return testing_data, nil
}

// Train will train the model, returning the best weights and the total number
// of successful and unsuccessful predictions found during training.
func Train(weight_vectors [][]float64, epochs int) ([][]float64, int, int, error) {
	verbose := true
	training_results, err := GetTrainingData(verbose)
	if err != nil {
		return nil, 0, 0, err
	}

	validation_results, err := GetValidationData(verbose)
	if err != nil {
		return nil, 0, 0, err
	}

	weights_after_each_epoch := [][][]float64{}
	successes_after_each_epoch := []int{}
	errors_after_each_epoch := []int{}

	learning_rate := 0.08 // eta η
	for epoch := 0; epoch < epochs; epoch++ {
		for _, row := range training_results {
			class_label := int(row[10])
			features := row[:10]

			logits := make([]float64, len(weight_vectors))
			for i, weights := range weight_vectors {
				logits[i] = helpers.DotProduct(weights, features)
			}

			predicted_label := helpers.ArgMax(logits)

			if predicted_label != class_label {
				adjusted_features := helpers.Multiply(features, learning_rate)
				weight_vectors[predicted_label] = helpers.SubtractVectors(weight_vectors[predicted_label], adjusted_features)
				weight_vectors[class_label] = helpers.AddVectors(weight_vectors[class_label], adjusted_features)
			}
		}

		successes, errors := Validate(weight_vectors, validation_results)

		weights_after_each_epoch = append(weights_after_each_epoch, weight_vectors)
		successes_after_each_epoch = append(successes_after_each_epoch, successes)
		errors_after_each_epoch = append(errors_after_each_epoch, errors)
	}

	best_epoch := helpers.ArgMin(errors_after_each_epoch)
	best_weights := weights_after_each_epoch[best_epoch]

	total_successes, total_errors := 0, 0
	for epoch := 0; epoch < epochs; epoch++ {
		total_successes += successes_after_each_epoch[epoch]
		total_errors += errors_after_each_epoch[epoch]
	}

	return best_weights, total_successes, total_errors, nil
}

// Validate will validate the provided weight vectors against the provided
// validation results, returning the total number of successful and unsuccessful
// predictions.
func Validate(weight_vectors [][]float64, validation_results [][]float64) (int, int) {
	successes, errors := 0, 0

	for _, row := range validation_results {
		class_label := int(row[10])
		features := row[:10]

		logits := make([]float64, len(weight_vectors))
		for i, weights := range weight_vectors {
			logits[i] = helpers.DotProduct(weights, features)
		}

		predicted_label := helpers.ArgMax(logits)

		if predicted_label == class_label {
			successes++
		} else {
			errors++
		}
	}

	return successes, errors
}

// GetPredictions will return a slice of ints representing the predicted labels
// for the provided file using the provided weight vectors.
func GetPredictions(file string, weight_vectors [][]float64) []int {
	testing_results, err := GetTestingData(file, false)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	predictions := make([]int, len(testing_results))

	for i, row := range testing_results {
		features := row[:10]

		logits := make([]float64, len(weight_vectors))
		for i, weights := range weight_vectors {
			logits[i] = helpers.DotProduct(weights, features)
		}

		predicted_label := helpers.ArgMax(logits)
		predictions[i] = predicted_label
	}

	return predictions
}

package main

import (
	"fmt"
	"os"

	"project04_perceptron/go_rewrite/display"
	"project04_perceptron/go_rewrite/helpers"
	"project04_perceptron/go_rewrite/model"
)

func main() {
	epochs, init_weights := 100, helpers.GetRandomWeights(10, 10)

	weights, num_successes, num_errors, err := model.Train(init_weights, epochs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	display.Stats(weights, num_successes, num_errors, epochs)

	test_file := "input_files/testing_data/unlabeled_digits.csv"
	predicted_labels := model.GetPredictions(test_file, weights)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Predicted Labels:")
	fmt.Println(predicted_labels)
}

package main

// Modelo main
// Incluir todos os packages

import (
	"fmt"

	"./Files"

	"./Forest"
)

func generateMetrics(y []int, ypred []int, nLabels int) {
	var total, right float64 = 0.0, 0.0

	classTotal := make([]float32, nLabels)
	classRight := make([]float32, nLabels)

	for i := 0; i < len(y); i++ {
		if y[i] == ypred[i] {
			right++
			classRight[y[i]]++
		}

		total++
		classTotal[y[i]]++
	}

	for i := 0; i < nLabels; i++ {
		fmt.Printf("Label %d => Accuracy = %.2f%%\n", i, 100*classRight[i]/classTotal[i])
	}

	fmt.Printf("Average accuracy: %.2f%%\n", 100*right/total)

}

func main() {
	//inputs := make([][]interface{}, 0)

	// Loading Dataset

	inputs := Files.ReadFile("./data/mnist_test.csv")

	fmt.Printf("%d\n", len(inputs))

	/*
		for _, line := range inputs {
			for _, val := range line {
				fmt.Printf("%f ", val)
			}
			fmt.Println()
		}*/

	// --- Splitting X and Y data

	X := make([][]float64, len(inputs))
	Y := make([]int, len(inputs))

	for i := 0; i < len(inputs); i++ {
		Y[i] = int(inputs[i][0])

		X[i] = inputs[i][1:len(inputs[0])]
	}

	fmt.Printf("\n---\n")

	rf := Forest.CreateRFClassifier(100, 20, 10)

	Forest.FitRFClassifier(rf, X, Y, len(X), len(X[0]))

	yPred := Forest.PredRFCLassifier(rf, X)

	for i := 0; i < 10; i++ {
		fmt.Printf("Sample %d predicted as %d\n", i, yPred[i])
	}

	fmt.Printf("\n---\n")

	generateMetrics(Y, yPred, 10)

}

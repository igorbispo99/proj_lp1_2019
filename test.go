package main

// Modelo main
// Incluir todos os packages

import (
	"fmt"

	"math/rand"

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

func splitTestTrain(x [][]float64, y []int, ratio float64) ([][]float64, []int, [][]float64, []int) {

	nSamples := int(float64(len(x)) * ratio)

	xTrain := make([][]float64, nSamples)
	yTrain := make([]int, nSamples)

	xTest := make([][]float64, len(x)-nSamples)
	yTest := make([]int, len(x)-nSamples)

	// Shuffling indexes vector
	idxRand := rand.Perm(len(x))

	for i := 0; i < nSamples; i++ {
		xTrain[i] = x[idxRand[i]]
		yTrain[i] = y[idxRand[i]]
	}

	for i := nSamples; i < len(x); i++ {
		xTest[i-nSamples] = x[idxRand[i]]
		yTest[i-nSamples] = y[idxRand[i]]
	}

	return xTrain, yTrain, xTest, yTest
}

func main() {
	// Loading Dataset

	inputs := Files.ReadFile("./data/mnist_test.csv")

	fmt.Printf("%d\n", len(inputs))

	// --- Splitting X and Y data

	X := make([][]float64, len(inputs))
	Y := make([]int, len(inputs))

	for i := 0; i < len(inputs); i++ {
		Y[i] = int(inputs[i][0])

		X[i] = inputs[i][1:len(inputs[0])]
	}

	fmt.Printf("\n---\n")

	// -- Splitting train and test data

	xTrain, yTrain, xTest, yTest := splitTestTrain(X, Y, 0.33)

	// -- Instantiating RF Classifier

	rf := Forest.CreateRFClassifier(100, 10, 10)

	// -- Fitting the data

	Forest.FitRFClassifier(rf, xTrain, yTrain, 1000, 28)

	// -- Predicting
	yPred := Forest.PredRFCLassifier(rf, xTest)

	for i := 0; i < 10; i++ {
		fmt.Printf("Sample %d predicted as %d\n", i, yPred[i])
	}

	fmt.Printf("\n---\n")

	generateMetrics(yTest, yPred, 10)

}

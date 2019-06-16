package main

// Modelo main
// Incluir todos os packages

import (
	"fmt"

	"./Files"

	"./Forest"
)

func main() {
	//inputs := make([][]interface{}, 0)

	// Loading Dataset

	inputs := Files.ReadFile("./data/test.csv")

	fmt.Printf("%d\n", len(inputs))

	/*
		for _, line := range inputs {
			for _, val := range line {
				fmt.Printf("%f ", val)
			}
			fmt.Println()
		}*/

	// --- Splitting X and Y data

	X := make([][]interface{}, len(inputs))
	Y := make([]int, len(inputs))

	for i := 0; i < len(inputs); i++ {

		y, ok := inputs[i][0].(int)

		if !ok {
			Y[i] = y
		} else {
			fmt.Printf("Could not parse label, aborting\n")

			return
		}

		X[i] = inputs[i][1:len(inputs[0])]
	}

	rf := Forest.CreateRFClassifier(300000, 4)

	Forest.FitRFClassifier(rf, X, Y, len(X), len(X[0]))

	yPred := Forest.PredRFCLassifier(rf, X) // TODO terminar a função interna Forest.Predict

	fmt.Printf("Sample 0 predicted as %d\n", yPred[0])

}

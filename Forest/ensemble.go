package Forest

import (
	"fmt"
	"sync"
)

type RFClassifier struct {
	forest   []Tree
	nTrees   int
	nThreads int
}

func CreateRFClassifier(nTrees int, nThreads int) *RFClassifier {

	rfClass := &RFClassifier{}

	rfClass.forest = make([]Tree, nTrees)

	rfClass.nTrees = nTrees

	rfClass.nThreads = nThreads

	return rfClass
}

func FitRFClassifier(rf *RFClassifier, inputs [][]interface{}, labels []int, nSamples int, nFeatures int) {
	// Initializing and Training trees

	fmt.Printf("Fitting data w/ %d trees and %d threads...\n", rf.nTrees, rf.nThreads)

	wg := new(sync.WaitGroup)

	for i := 0; i < rf.nTrees; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			tree := NewTree(inputs, labels, nSamples, nFeatures)
			rf.forest[i] = *tree

		}(i)
	}

	go func() {
		wg.Wait()
	}()

	fmt.Printf("Done...\n")

}

func PredRFCLassifier(rf *RFClassifier, inputs [][]interface{}) []int {

	preds := make([][]int, rf.nTrees)

	wg := new(sync.WaitGroup)

	for i := 0; i < rf.nTrees; i++ {
		wg.Add(1)

		go func(i int) {
			// TODO Função Forest.Predict não existe ainda
			// TODO Descomentar a linha seguinte quando Predict existir

			// preds[i] = Forest.Predict(rf.forest[i], inputs)
		}(i)
	}

	go func() {
		wg.Wait()
	}()

	// Returning most frequent value to each sample

	yPredicted := make([]int, len(inputs))

	nLabels := len(inputs[0])

	for j := 0; j < len(inputs); j++ {

		wg.Add(1)

		go func(j int) {

			// (Probably) This loop isnt multi-thread"ible"
			for i := 0; i < rf.nTrees; i++ {
				preds[preds[i][j]%nLabels][j] += nLabels
			}

			max := preds[0][j]
			yPredicted[j] = 0

			// Maybe this loop is "multi-thread"ible"
			for i := 1; i < rf.nTrees; i++ {

				if preds[i][j] > max {
					max = preds[i][j]
					yPredicted[j] = j
				}

			}

		}(j)

	}

	go func() {
		wg.Wait()
	}()

	return yPredicted
}

/// Rascunho

//func main() {
//	rf := createRFClassifier(100)
//	fmt.Print(rf.forest)
//}

/// endOf Rascunho

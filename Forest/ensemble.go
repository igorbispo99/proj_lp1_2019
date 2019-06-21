package Forest

import (
	"fmt"
	"sync"
)

type RFClassifier struct {
	forest   []*Tree
	nTrees   int
	nLabels  int
	maxDepth int
}

func CreateRFClassifier(nTrees int, maxDepth int, nLabels int) *RFClassifier {

	rfClass := &RFClassifier{}

	rfClass.forest = make([]*Tree, nTrees)

	rfClass.nTrees = nTrees

	//rfClass.nThreads = nThreads

	rfClass.nLabels = nLabels

	rfClass.maxDepth = maxDepth

	return rfClass
}

func FitRFClassifier(rf *RFClassifier, inputs [][]float64, labels []int, nSamples int, nFeatures int) {
	// Initializing and Training trees

	fmt.Printf("Fitting data w/ %d trees, %d features/samples and %d samples/tree...\n", rf.nTrees, nFeatures, nSamples)

	wg_ := new(sync.WaitGroup)

	for i := 0; i < rf.nTrees; i++ {
		wg_.Add(1) 

		go func(i int) {
			rf.forest[i] = NewTree(inputs, labels, nSamples, nFeatures, rf.maxDepth)
			wg_.Done()
		}(i)

	}
	wg_.Wait()

	fmt.Printf("Done...\n")
}

func PredRFCLassifier(rf *RFClassifier, inputs [][]float64) []int {

	preds := make([][]int, rf.nTrees)

	wg := new(sync.WaitGroup)

	for i := 0; i < rf.nTrees; i++ {
		wg.Add(1)

		go func(i int) {
			// TODO Função Forest.Predict não existe ainda
			// TODO Descomentar a linha seguinte quando Predict existir

			preds[i] = PredictTree(rf.forest[i], inputs)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Returning most frequent value to each sample

	yPredicted := make([]int, len(inputs))

	for j := 0; j < len(inputs); j++ {

		wg.Add(1)

		go func(j int) {

			yHist := make([]int, rf.nLabels)

			// (Probably) This loop isnt multi-thread"ible"
			for i := 0; i < rf.nTrees; i++ {
				yHist[preds[i][j]] += 1
			}

			max := yHist[0]
			maxIdx := 0

			// Maybe this loop is "multi-thread"ible"
			for i := 1; i < rf.nLabels; i++ {

				if yHist[i] > max {
					max = yHist[i]
					maxIdx = i
				}
			}

			yPredicted[j] = maxIdx

			wg.Done()
		}(j)

	}

	wg.Wait()

	return yPredicted
}

/// Rascunho

//func main() {
//	rf := createRFClassifier(100)
//	fmt.Print(rf.forest)
//}

/// endOf Rascunho

package forest

import( //"fmt"
		//"math"
		"math/rand"
		//classifier

)

// Tree Struct ____________________________________________________________
type Tree struct{
	root *Node

	features  []int 
	nFeatures int

	maxDepth int 
	nNodes int
	nClasses int
}

func newTree(inputs [][]interface{}, labels []int, nSamples int, nFeatures int) *Tree {
	samplesLabels := make([]int, nSamples)
	samples := make([][]interface{}, nSamples)

	for i:=0; i<nSamples; i++{
		j := int(rand.Float64()*float64(len(inputs))) // retorna o inteiro do float aleatório gerado entre [0.0, 1.0] * len(inputs) ou seja entre [0.0, len(inputs)]
		samples[i] = inputs[j]
		samplesLabels[i] = labels[j]
	}

	tree := &Tree{}
	tree.root = buildTree(samples, samplesLabels, nFeatures)

	return tree
}

func buildTree(samples [][]interface{}, samplesLabels []int, nFeatures int) *Node{
	nColumns := len(samples[0]) 	// amostras
	nRows := nFeatures				// características
	selectedColumns := 	randomColumns(nColumns, nFeatures)			

	bestGain := 0.0
	bestPartL := make([]int, 0, len(samples))
	bestPartR := make([]int, 0, len(samples))
	bestTotalL := 0 
	bestTotalR := 0
	var bestThreshold interface{}
	var bestCol int
	var best_column_type string

	total := len(samplesLabels)
	pMap := make(map[int]float64)   // pMap : probability map
	for i:=0; i<total; i++{  
		pMap[samplesLabels[i]] += 1	
    }

	//entropy := Entropy(pMap, total) // Criterio de impureza escolhido. Testar com Gini

	for _,col := range selectedColumns{
		// encontra melhor ponto de corte dado a região atual do espaço de acordo com critério de impureza escolhido
		gain, threshold, totalL, totalR := cutPoint(samples, col, samplesLabels, entropy)

		if gain >= bestGain{
			bestGain = gain
			bestThreshold = threshold
			bestCol = col
			bestTotalL = totalL
			bestTotalR = totalR
		}
	}

	if bestGain > 0 && bestTotalL > 0 && bestTotalR > 0 {  
		node := &Node{}
		node.value = bestThreshold
		node.columnNumber = bestCol
		// aplicando o limiar para particionar o espaço:
		//splitSamples(samples, bestCol, bestThreshold, &bestPartL, &bestPartR)

		// contruindo árvore recursivamente para os sub-espaço da direita e da esquerda.
		//node.left = buildTree(getSamples(samples, bestPartL), getLabels(samplesLabels, bestPartL),  nFeatures)
		//node.right = buildTree(getSamples(samples, bestPartR), getLabels(samplesLabels, bestPartR),  nFeatures)
		return node
	}
	
	return leafNode(samplesLabels)
}

// Node Struct _________________________________________________________
type Node struct{
	value interface{}
	labels map[int]int
	//decision int
	//threshold int
	columnNumber int 
	left *Node	
	right *Node
}

func leafNode(labels []int) *Node{
	counter := make(map[int]int)
	for _,label := range labels{
		counter[label] += 1
	}

	node := &Node{}
	node.labels = counter
	return node
}
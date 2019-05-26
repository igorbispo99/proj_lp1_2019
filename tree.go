package forest

import( "fmt"
		"math"
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

//func buildTree(samples [][]interface{}, samplesLabels []int, nFeatures int) *Node/{	
//}

// Node Struct _________________________________________________________
type Node struct{
	value interface{}
	labels []int
	//decision int
	//threshold int
	columnNumber int 
	left *Node	
	right *Node
}


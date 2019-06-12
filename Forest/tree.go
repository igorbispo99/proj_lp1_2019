package Forest

import( //"fmt"
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

func NewTree(inputs [][]interface{}, labels []int, nSamples int, nFeatures int) *Tree {
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
	nColumns := len(samples[1]) 	// amostras
	//nRows := nFeatures				// características
	selectedColumns := 	randomColumns(nColumns, nFeatures)			

	bestGain := 0.0
	bestPartL := make([]int, 0, len(samples))
	bestPartR := make([]int, 0, len(samples))
	bestTotalL := 0 
	bestTotalR := 0
	var bestThreshold interface{}
	var bestCol int

	total := len(samplesLabels)
	pMap := make(map[int]float64)   // pMap : probability map
	for i:=0; i<total; i++{  
		pMap[samplesLabels[i]] += 1	
    }

	entropy := Entropy(pMap, total) // Criterio de impureza escolhido. Testar com Gini

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
		splitSamples(samples, bestCol, bestThreshold, &bestPartL, &bestPartR)

		// contruindo árvore recursivamente para os sub-espaço da direita e da esquerda.
		node.left = buildTree(getSamples(samples, bestPartL), getLabels(samplesLabels, bestPartL),  nFeatures)
		node.right = buildTree(getSamples(samples, bestPartR), getLabels(samplesLabels, bestPartR),  nFeatures)
		return node
	}
	
	return leafNode(samplesLabels)
}

func randomColumns(nColumns int, nFeatures int) []int{
	tmp := make([]int, nColumns)
	for i:=0; i<nColumns; i++{ tmp[i]=i }
	for i:=0; i<nFeatures; i++{
		j := i + int(rand.Float64()*float64(nColumns-i))
		tmp[i],tmp[j] = tmp[j],tmp[i]
	}

	return tmp[:nFeatures]
}

func Entropy(p_vec map[int]float64, total int) float64 {
	entropy := 0.0

	for idx,_ := range p_vec{ p_vec[idx] = p_vec[idx] / float64(total) } // probability
	for _,p := range p_vec{  entropy -= p*math.Log2(p)  } 

	return entropy
}

func Gini(p_vec map[string]float64) float64 {   // VERIFICAR
	total := 0.0
	impure := 0.0

	for _,v := range p_vec{  total += v  }
	for k,_ := range p_vec{  p_vec[k] = p_vec[k] / total  }

	
	for k1, v1 := range p_vec{
		for k2, v2 := range p_vec{
			if k1 != k2{
				impure += v1*v2
			}
		} 
	}
	return impure
}

func cutPoint(samples [][]interface{}, c int, samplesLabels []int, currentEntropy float64) (float64,interface{}, int, int){
	var bestThreshold interface{}
	bestGain := 0.0
	totalR := 0
	totalL := 0

	uniqValues := make(map[interface{}]int)  // valores ordenados únicos da característica c (coluna c)
	for i:=0;i<len(samples);i++{			 // map só tem os índices atribuídos de 1, outros não existem
		uniqValues[samples[i][c]] = 1		// atribui 1 só para fazer existir o indice
	}

	for value,_ := range uniqValues{		// valor da característica = indice do mapa
		mapL := make(map[int]float64)
		mapR := make(map[int]float64)
		currTotalL := 0
		currTotalR := 0 
		for j:=0; j<len(samples); j++{
			if samples[j][c].(float64) <= value.(float64){
				currTotalL += 1
				mapL[samplesLabels[j]] += 1.0
			}else{
				currTotalR += 1
				mapR[samplesLabels[j]] += 1.0
			}
		}
		

		p1 := float64(currTotalR) / float64(len(samples))
		p2 := float64(currTotalL) / float64(len(samples))

		newEntropy := p1*Entropy(mapR, currTotalR) + p2*Entropy(mapL, currTotalL)
		entropyGain := currentEntropy - newEntropy
		
		if entropyGain >= bestGain{
			bestGain = entropyGain
			bestThreshold = value
			totalL = currTotalL
			totalR = currTotalR
		}
	}

	return bestGain, bestThreshold, totalL, totalR
}

func splitSamples(samples [][]interface{}, c int, value interface{}, partL *[]int, partR *[]int){
	for j:=0; j<len(samples); j++{
		if samples[j][c].(float64) <= value.(float64){
			*partL = append(*partL,j)
		
		}else{
			*partR = append(*partR,j)
		}
	}
	
}

func getSamples(samples [][]interface{}, idxPart []int)  [][]interface{} {
	samplesPart := make([][]interface{}, len(idxPart))
	for i:=0; i<len(idxPart); i++{
		samplesPart[i] = samples[idxPart[i]]
	}
	return samplesPart
}


func getLabels(samples []int, idxPart []int ) []int{
	labelsPart := make([]int,len(idxPart))
	for i:=0; i<len(idxPart); i++{
		labelsPart[i] = samples[idxPart[i]]
	}
	return labelsPart
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
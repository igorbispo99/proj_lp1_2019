package Forest

import ( //"fmt"
	"math"
	"math/rand"
	"sync"
)

// Tree Struct ____________________________________________________________
type Tree struct {
	root  *Node
	depth int
}

// Retorna uma nova árvore com nSamples (número de amostras) aleatórias do input.
func NewTree(input [][]interface{}, labels []int, nSamples int, nSelectedFeatures int, maxDepth int) *Tree {
	samplesLabels := make([]int, nSamples)
	samples := make([][]interface{}, nSamples)

	wg := sync.WaitGroup{}

	for i := 0; i < nSamples; i++ {
		wg.Add(1)
		go func(i int) {
			j := int(rand.Float64() * float64(len(input))) // retorna o inteiro do float aleatório gerado entre [0.0, 1.0] * len(input) ou seja entre [0.0, len(input)]
			samples[i] = input[j]
			samplesLabels[i] = labels[j]
			wg.Done() // decrementa o contador WaitGroup por um.
		}(i)
	}
	wg.Wait() // Aguarde até que o contador WaitGroup seja zero. Espera todas as goroutines terminarem
	tree := &Tree{}
	tree.depth = 1
	tree.root = buildTree(samples, samplesLabels, nSelectedFeatures, maxDepth, tree)

	return tree
}

// Cria a árvore recursivamente
// - Seleciona algumas Características, - Calcula a Entropia, - encontra melhor ponto de corte (Característica de Corte),
// - Separa os Rótulos, - Calcula sub-árvores à esquerda e à direita
func buildTree(samples [][]interface{}, samplesLabels []int, nSelectedFeatures int, maxDepth int, tree *Tree) *Node {

	if tree.depth != maxDepth {
		tree.depth += 1
		nFeatures := len(samples[1])                                     // Features
		selectedFeatures := randomFeatures(nFeatures, nSelectedFeatures) // características selecionadas para o cutPoint

		bestGain := 0.0
		bestPartL := make([]int, 0, len(samples))
		bestPartR := make([]int, 0, len(samples))
		bestTotalL := 0
		bestTotalR := 0
		var bestThreshold interface{}
		var bestThresholdIdx int

		total := len(samplesLabels)
		qtdMap := make(map[int]float64) // qtdMap : mapa das quatidades de amostras por Rótulo
		for i := 0; i < total; i++ {
			qtdMap[samplesLabels[i]] += 1
		}

		entropy := Entropy(qtdMap, total) // Criterio de impureza escolhido. Testar com Gini

		// encontra melhor ponto de corte dado a região atual do espaço de acordo com critério de impureza escolhido
		for _, col := range selectedFeatures {
			gain, threshold, totalL, totalR := cutPoint(samples, col, samplesLabels, entropy)

			if gain >= bestGain {
				bestGain = gain
				bestThreshold = threshold
				bestThresholdIdx = col
				bestTotalL = totalL
				bestTotalR = totalR
			}
		}

		if bestGain > 0 && bestTotalL > 0 && bestTotalR > 0 {
			node := &Node{}
			node.threshold = bestThreshold
			node.thredholdIdx = bestThresholdIdx
			// aplicando o limiar para particionar o espaço:
			splitSpace(samples, bestThresholdIdx, bestThreshold, &bestPartL, &bestPartR)

			// contruindo árvore recursivamente para os sub-espaço da direita e da esquerda.

			/* //contruindo Sub-espaços da direita e da esquerda em paralelo
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func (){
				node.left = buildTree(getSamples(samples, bestPartL), getLabels(samplesLabels, bestPartL), nFeatures, maxDepth, tree)
				wg.Done()
			}
			go func(){
				node.right = buildTree(getSamples(samples, bestPartR), getLabels(samplesLabels, bestPartR), nFeatures, maxDepth, tree)
				wg.Done()
			}

			wg.Wait() */

			node.left = buildTree(getSamples(samples, bestPartL), getLabels(samplesLabels, bestPartL), nFeatures, maxDepth, tree)
			node.right = buildTree(getSamples(samples, bestPartR), getLabels(samplesLabels, bestPartR), nFeatures, maxDepth, tree)
			return node
		} else {
			return leafNode(samplesLabels)
		}
	}

	return leafNode(samplesLabels)
}

// Retorna uma vetor com nSelectedFeatures índices das Características. Para isso, primeiro embaralha as Características.
func randomFeatures(nFeatures int, nSelectedFeatures int) []int {
	tmp := make([]int, nFeatures)
	for i := 0; i < nFeatures; i++ {
		tmp[i] = i
	}
	for i := 0; i < nSelectedFeatures; i++ {
		j := i + int(rand.Float64()*float64(nFeatures-i)) // j: número aleatório entre 0 e nFeatures
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	return tmp[:nSelectedFeatures]
}

// Input: Árvore de decisão e amostras
// Output: Vetor de classificação das amostras
func predictTree(tree Tree, input [][]interface{}) []int {
	predictions := make([]int, len(input))

	// Verificar e classificar cada amostra
	for i, sample := range input {

		// Inicializa o nó atual que
		// percorrerá a árvore
		currentNode := tree.root
		labelsLength := len(currentNode.labels)

		// Percorre a árvore até o nó folha
		// Seguindo à esquerda se o valor da amostra
		// é menor que o threshold, caso contrário à direita
		for labelsLength == 0 {
			feature := currentNode.thredholdIdx
			threshold := currentNode.threshold.(float32)

			if sample[feature].(float32) < threshold {
				currentNode = currentNode.left
			} else {
				currentNode = currentNode.right
			}
		}

		// Encontrar o índice do label de maior ocorrência
		predictions[i] = maxOccurrenceLabel(currentNode.labels)

	}

	return predictions
}

// Input: lista de ocorrência de labels exemplos
// Output: o label de maior ocorrência
func maxOccurrenceLabel(listLabels map[int]int) int {
	var maxLabel = 0

	for i := 1; i < len(listLabels); i++ {
		if listLabels[maxLabel] < listLabels[i] {
			maxLabel = i
		}
	}

	return maxLabel
}

func Entropy(p_vec map[int]float64, total int) float64 {
	entropy := 0.0

	for idx, _ := range p_vec {
		p_vec[idx] = p_vec[idx] / float64(total)
	} // probability
	for _, p := range p_vec {
		entropy -= p * math.Log2(p)
	}

	return entropy
}

func Gini(p_vec map[string]float64) float64 { // VERIFICAR
	total := 0.0
	impure := 0.0

	for _, v := range p_vec {
		total += v
	}
	for k, _ := range p_vec {
		p_vec[k] = p_vec[k] / total
	}

	for k1, v1 := range p_vec {
		for k2, v2 := range p_vec {
			if k1 != k2 {
				impure += v1 * v2
			}
		}
	}
	return impure
}

/* Cada divisão do espaço é representada por um nó na árvore de decisão. A primeira divisão (nós raiz da árvore) leva em consideração todos os exemplos (separados aleatóriamente) do espaço ao encontrar o ponto de corte que maximiza a pureza, de acordo com algum critério de impureza, das sub-regiões resultantes.
Para encontrar melhor ponto de corte, são testados todos os possíveis, ou seja, para cada atributo e valores possíveis calcula-se o ganho de informação (quão pura a divisão torna o espaço) para cada um dos pontos de corte candidatos. Após essa etapa, é escolhido o candidato com maior ganho de informação para ser o ponto de corte do nó em questão.
*/
func cutPoint(samples [][]interface{}, c int, samplesLabels []int, currentEntropy float64) (float64, interface{}, int, int) {
	var bestThreshold interface{}
	bestGain := 0.0
	totalR := 0
	totalL := 0

	uniqValues := make(map[interface{}]int) // valores ordenados únicos da característica c (coluna c)
	for i := 0; i < len(samples); i++ {     // map só tem os índices atribuídos de 1, outros não existem
		uniqValues[samples[i][c]] = 1 // atribui 1 só para fazer existir o indice
	}

	for value, _ := range uniqValues { // valor da característica = indice do mapa
		mapL := make(map[int]float64)
		mapR := make(map[int]float64)
		currTotalL := 0
		currTotalR := 0
		for j := 0; j < len(samples); j++ {
			if samples[j][c].(float64) <= value.(float64) {
				currTotalL += 1
				mapL[samplesLabels[j]] += 1.0
			} else {
				currTotalR += 1
				mapR[samplesLabels[j]] += 1.0
			}
		}

		p1 := float64(currTotalR) / float64(len(samples))
		p2 := float64(currTotalL) / float64(len(samples))

		newEntropy := p1*Entropy(mapR, currTotalR) + p2*Entropy(mapL, currTotalL)
		entropyGain := currentEntropy - newEntropy

		if entropyGain >= bestGain {
			bestGain = entropyGain
			bestThreshold = value
			totalL = currTotalL
			totalR = currTotalR
		}
	}

	return bestGain, bestThreshold, totalL, totalR
}

// Cria os sub-espaço da direita e da esquerda (separa o espaço das amostras). partL e partR são vetores de indices
func splitSpace(samples [][]interface{}, c int, value interface{}, partL *[]int, partR *[]int) {
	for j := 0; j < len(samples); j++ {
		if samples[j][c].(float64) <= value.(float64) {
			*partL = append(*partL, j)

		} else {
			*partR = append(*partR, j)
		}
	}

}

// Função auxiliar para pegar as amotras de um dado sub-espaço
func getSamples(samples [][]interface{}, idxPart []int) [][]interface{} {
	samplesPart := make([][]interface{}, len(idxPart))
	for i := 0; i < len(idxPart); i++ {
		samplesPart[i] = samples[idxPart[i]]
	}
	return samplesPart
}

// Função auxiliar para pegar os Rótulos de um dado sub-espaço
func getLabels(samples []int, idxPart []int) []int {
	labelsPart := make([]int, len(idxPart))
	for i := 0; i < len(idxPart); i++ {
		labelsPart[i] = samples[idxPart[i]]
	}
	return labelsPart
}

// Node Struct _________________________________________________________
type Node struct {
	threshold    interface{}
	thredholdIdx int
	labels       map[int]int // quantidade de exemplos por classe nesse nó folha
	left         *Node
	right        *Node
}

// Retorna Nó folha
func leafNode(labels []int) *Node {
	counter := make(map[int]int) //quantidade de exemplos por classe nesse nó folha
	for _, label := range labels {
		counter[label] += 1
	}

	node := &Node{}
	node.labels = counter
	return node
}

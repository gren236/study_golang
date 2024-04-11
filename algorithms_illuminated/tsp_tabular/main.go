package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	x, y float64
}

func parseTestFile() []node {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/tsp_tabular/tsp.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")

	size, _ := strconv.Atoi(inputStrings[0])

	res := make([]node, size)
	for i, inputString := range inputStrings[1 : len(inputStrings)-1] {
		row := strings.Split(inputString, " ")
		x, _ := strconv.ParseFloat(row[0], 64)
		y, _ := strconv.ParseFloat(row[1], 64)

		res[i] = node{x, y}
	}

	return res
}

func calcDist(a, b node) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

// solveTSP uses Bellman-Held-Karp algorithm to solve Travelling Salesman Problem (dynamic)
func solveTSP(nodes []node) float64 {
	n := len(nodes)
	s := binSet{length: uint(len(nodes))}

	// Generate 2D array for sub-problems
	a := make([][]float64, int(math.Pow(2, float64(n)))) // a[set_enc][node_label]=dist
	for i := range a {
		a[i] = make([]float64, n)
	}

	// Base cases
	for j := 1; j < n; j++ {
		s2 := s.encode([]int{0, j})
		d := calcDist(nodes[0], nodes[j])
		a[s2][j] = d
	}

	// Go through all sub-problems
	for m := 3; m <= n; m++ {
		c := combin.NewCombinationGenerator(n, m)
		subset := make([]int, m)

		fmt.Printf("sub-problem size: %d\n", m)

		for c.Next() {
			c.Combination(subset)

			sSub := s.encode(subset)

			for _, j := range subset {
				// j ∈ S - {0}
				if j == 0 {
					continue
				}

				sSubMinusJ := s.remove(sSub, j)

				minDist := math.Inf(1)
				for _, k := range subset {
					// k ∈ S; k != 0,j
					if k == 0 || k == j {
						continue
					}

					distKJ := a[sSubMinusJ][k] + calcDist(nodes[k], nodes[j])

					if distKJ < minDist {
						minDist = distKJ
					}
				}

				a[sSub][j] = minDist
			}
		}
	}

	// Calculate most optimal biggest sub-problem
	res := math.Inf(1)
	for j := 1; j < n; j++ {
		distJ0 := a[int(math.Pow(2, float64(n)))-1][j] + calcDist(nodes[j], nodes[0])

		if distJ0 < res {
			res = distJ0
		}
	}

	return res
}

func main() {
	nodes := parseTestFile()

	res := solveTSP(nodes)

	fmt.Println(res)
}

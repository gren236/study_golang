package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	from, to int
}

// parseListFile returns each edge mapped to its weight and number of vertices
func parseListFile(path string) (map[edge]int, int) {
	//inputRaw, _ := os.ReadFile("./algorithms_illuminated/all_pair_shortest_path/input_random_10_8.txt") // -41
	//inputRaw, _ := os.ReadFile("./algorithms_illuminated/all_pair_shortest_path/input_random_12_8.txt") // -234
	inputRaw, _ := os.ReadFile(path) // -234
	inputStrings := strings.Split(string(inputRaw), "\n")
	headRow := strings.Split(inputStrings[0], " ")

	lengthV, _ := strconv.Atoi(headRow[0])
	lengthE, _ := strconv.Atoi(headRow[1])
	res := make(map[edge]int, lengthE)

	inputStrings = inputStrings[1 : len(inputStrings)-1]
	for _, s := range inputStrings {
		inputRow := strings.Split(s, " ")

		from, _ := strconv.Atoi(inputRow[0])
		to, _ := strconv.Atoi(inputRow[1])
		w, _ := strconv.Atoi(inputRow[2])

		e := edge{from, to}

		if wFound, ok := res[e]; ok {
			res[e] = min(w, wFound)
		} else {
			res[e] = w
		}
	}

	return res, lengthV
}

// floydWarshall returns shortest path distance for each (u, v) vertex pair or error if negative cost cycle found.
func floydWarshall(edges map[edge]int, n int) ([][]float64, error) {
	// Initiate the sub-problem array: [k-prefix][u][v] = memoized weight
	subp := make([][][]float64, n+1)
	for i := range subp {
		subp[i] = make([][]float64, n)

		for j := range subp[i] {
			subp[i][j] = make([]float64, n)
		}
	}

	// Base cases k=0
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if u == v {
				subp[0][u][v] = 0

				continue
			}

			if w, ok := edges[edge{u + 1, v + 1}]; ok {
				subp[0][u][v] = float64(w)

				continue
			}

			subp[0][u][v] = math.Inf(1)
		}
	}

	// Solve sub-problems iteratively
	for k := 1; k <= n; k++ {
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				subp[k][u][v] = min(subp[k-1][u][v], subp[k-1][u][k-1]+subp[k-1][k-1][v])
			}
		}
	}

	// Check for negative cost cycles
	for v := 0; v < n; v++ {
		if subp[n][v][v] < 0 {
			return nil, errors.New("negative-cost cycle found")
		}
	}

	// Form resulting slice
	res := make([][]float64, n)
	for u := range res {
		res[u] = make([]float64, n)

		for v := range res[u] {
			res[u][v] = subp[n][u][v]
		}
	}

	return res, nil
}

func getMinFromResults(res [][]float64) int {
	m := res[0][0]

	for u := range res {
		for _, d := range res[u] {
			if d < m {
				m = d
			}
		}
	}

	return int(m)
}

func main() {
	paths := []string{
		"./algorithms_illuminated/all_pair_shortest_path/g1.txt",
		"./algorithms_illuminated/all_pair_shortest_path/g2.txt",
		"./algorithms_illuminated/all_pair_shortest_path/g3.txt",
	}

	for i, path := range paths {
		edges, nVertices := parseListFile(path)

		var output string

		res, err := floydWarshall(edges, nVertices)
		if err != nil {
			output = err.Error()
		} else {
			output = fmt.Sprintf("%d", getMinFromResults(res))
		}

		fmt.Printf("#%d result: %s\n", i, output)
	}
}

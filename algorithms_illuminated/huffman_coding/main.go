package main

import (
	"fmt"
	"github.com/gren236/study_golang/pkg/heap"
	"os"
	"strconv"
	"strings"
)

func parseListFile() []*Node {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/huffman_coding/huffman.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")

	length, _ := strconv.Atoi(inputStrings[0])
	resSlice := make([]*Node, length)

	inputStrings = inputStrings[1 : len(inputStrings)-1]
	for i, s := range inputStrings {
		inputRow := strings.Split(s, " ")

		w, _ := strconv.Atoi(inputRow[0])

		resSlice[i] = NewNode(w)
	}

	return resSlice
}

func main() {
	s := parseListFile()
	h := heap.NewHeapFromSlice(s)

	for h.Len() > 1 {
		t1 := h.ExtractMin()
		t2 := h.ExtractMin()

		tMerged := t1.Merge(t2)

		h.Insert(tMerged)
	}

	res := h.PeekMin()

	fmt.Printf("max: %v, min: %v\n", res.maxRank, res.minRank)
}

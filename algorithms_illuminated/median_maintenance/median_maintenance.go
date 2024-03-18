package main

import (
	"fmt"
	"github.com/gren236/study_golang/pkg/heap"
	"os"
	"strconv"
	"strings"
)

type intLess int

func (i intLess) Less(t intLess) bool {
	return i < t
}

type intGreater int

func (i intGreater) Less(t intGreater) bool {
	return i > t
}

func main() {
	// Parse input file
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/median_maintenance/Median.txt")
	inputStrings := strings.Split(string(inputRaw), "\r\n")
	inputStrings = inputStrings[:len(inputStrings)-1]

	hLow := heap.Container[intGreater]{}
	hHigh := heap.Container[intLess]{}

	// Handle first entry
	fv, _ := strconv.Atoi(inputStrings[0])
	hLow.Insert(intGreater(fv))

	mSum := fv

	// Handle each entry
	for i, str := range inputStrings[1:] {
		v, _ := strconv.Atoi(str)

		if int(hLow.PeekMin()) >= v {
			hLow.Insert(intGreater(v))
		} else {
			hHigh.Insert(intLess(v))
		}

		// Balance heaps
		if hLow.Len()-hHigh.Len() > 1 {
			hHigh.Insert(intLess(hLow.ExtractMin()))
		}

		if hHigh.Len()-hLow.Len() > 1 {
			hLow.Insert(intGreater(hHigh.ExtractMin()))
		}

		// Get median
		if hLow.Len() > hHigh.Len() {
			mSum += int(hLow.PeekMin())
			continue
		}

		if hLow.Len() < hHigh.Len() {
			mSum += int(hHigh.PeekMin())
			continue
		}

		if (i+2)%2 == 0 {
			mSum += int(hLow.PeekMin())
		} else {
			mSum += int(hHigh.PeekMin())
		}
	}

	fmt.Println(mSum % 10000)
}

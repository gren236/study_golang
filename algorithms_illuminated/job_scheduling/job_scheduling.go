package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type job struct {
	w, l int
}

type ByGreedDiff []job

func (gd ByGreedDiff) Len() int {
	return len(gd)
}

func (gd ByGreedDiff) Less(i, j int) bool {
	diffI := gd[i].w - gd[i].l
	diffJ := gd[j].w - gd[j].l

	if diffI == diffJ {
		return gd[i].w > gd[j].w
	}

	return diffI > diffJ
}

func (gd ByGreedDiff) Swap(i, j int) {
	gd[i], gd[j] = gd[j], gd[i]
}

type ByGreedRatio []job

func (gr ByGreedRatio) Len() int {
	return len(gr)
}

func (gr ByGreedRatio) Less(i, j int) bool {
	return (float64(gr[i].w) / float64(gr[i].l)) > (float64(gr[j].w) / float64(gr[j].l))
}

func (gr ByGreedRatio) Swap(i, j int) {
	gr[i], gr[j] = gr[j], gr[i]
}

func parseJobsFromFile() []job {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/job_scheduling/jobs.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")
	inputStrings = inputStrings[:len(inputStrings)-1]

	length, _ := strconv.Atoi(inputStrings[0])
	inputStrings = inputStrings[1:]

	res := make([]job, length)
	for i, s := range inputStrings {
		sArr := strings.Split(s, " ")

		w, _ := strconv.Atoi(sArr[0])
		l, _ := strconv.Atoi(sArr[1])

		res[i] = job{w, l}
	}

	return res
}

func getSumOfWeightedCompletionTimes(jobs []job) int {
	var compTime, sum int

	for _, j := range jobs {
		compTime += j.l

		sum += compTime * j.w
	}

	return sum
}

func main() {
	inJobsDiff := parseJobsFromFile()
	inJobsRatio := parseJobsFromFile()

	sort.Sort(ByGreedDiff(inJobsDiff))

	fmt.Println(getSumOfWeightedCompletionTimes(inJobsDiff))

	sort.Sort(ByGreedRatio(inJobsRatio))

	fmt.Println(getSumOfWeightedCompletionTimes(inJobsRatio))
}

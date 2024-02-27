package main

import (
	"fmt"
	"github.com/gren236/study_golang/pkg/hasht"
	"os"
	"strconv"
	"strings"
)

type tInt int

func (ti tInt) Bytes() []byte {
	return []byte(strconv.Itoa(int(ti)))
}

func parseIntFileToArray() []tInt {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/two_sum/2sum.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")
	inputStrings = inputStrings[:len(inputStrings)-1]

	res := make([]tInt, len(inputStrings))

	for i, s := range inputStrings {
		v, _ := strconv.Atoi(s)

		res[i] = tInt(v)
	}

	return res
}

func twoSumCustom(a []tInt, ht *hasht.Table[tInt, bool], t int) bool {
	for _, v := range a {
		diff := t - int(v)

		if diff == int(v) {
			continue
		}

		if _, ok := ht.Get(tInt(diff)); ok {
			return true
		}
	}

	return false
}

func twoSumMap(a []tInt, ht map[tInt]bool, t int) bool {
	for _, v := range a {
		diff := t - int(v)

		if diff == int(v) {
			continue
		}

		if _, ok := ht[tInt(diff)]; ok {
			return true
		}
	}

	return false
}

func main() {
	a := parseIntFileToArray()

	// Populate hash table and std map for comparison
	ht := hasht.New[tInt, bool]()
	mp := make(map[tInt]bool)

	for _, v := range a {
		ht.Insert(v, false)
		mp[v] = false
	}

	// count sums of t in range [-10000,10000]
	var cnt int

	for t := -10000; t <= 10000; t++ {
		if t%1000 == 0 {
			fmt.Printf("current t is: %d, and cnt=%d\n", t, cnt)
		}

		if twoSumCustom(a, ht, t) {
			cnt++
		}

		//if twoSumMap(a, mp, t) {
		//	cnt++
		//}
	}

	fmt.Printf("Done! cnt=%d", cnt)
}

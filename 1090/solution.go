package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	val int
	lab int
}

func main() {
	values := []int{9, 2, 8, 7, 6}
	labels := []int{0, 0, 5, 1, 1}
	numWanted := 3
	useLimit := 3
	fmt.Println(largestValsFromLabels(values, labels, numWanted, useLimit))
}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	usedLabels := make(map[int]int)
	pairedArr := make([]Pair, len(values))
	for i := 0; i < len(values); i++ {
		pairedArr = append(pairedArr, Pair{values[i], labels[i]})
	}

	sort.Slice(pairedArr, func(i, j int) bool {
		return pairedArr[i].val > pairedArr[j].val
	})

	retval := 0
	numsAdded := 0
	for i := 0; i < len(pairedArr); i++ {
		if numsAdded >= numWanted {
			break
		}
		if usedLabels[pairedArr[i].lab] < useLimit {
			retval += pairedArr[i].val
			usedLabels[pairedArr[i].lab]++
			numsAdded++
		}
	}

	return retval
}

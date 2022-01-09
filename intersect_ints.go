package go_facet_indexer

import "math"

func Intersect(input ...[]int) (res []int) {
	for i := range input {
		if len(input[i]) == 0 {
			return res
		}
	}
	iterators := make([]int, len(input))
	ends := make([]bool, len(input))
	for {
		minI := -1
		min := math.MaxInt
		curr := input[0][iterators[0]]
		eq := true
		for numInput, iterInput := range iterators {
			el := input[numInput][iterInput]
			if curr != el {
				eq = false
			}
			if !ends[numInput] {
				if el < min {
					min = el
					minI = numInput
				}
			}
		}
		if minI == -1 {
			break
		}
		if eq {
			res = append(res, curr)
			for numInput := range iterators {
				iterators[numInput]++
				if iterators[numInput] == len(input[numInput]) {
					ends[numInput] = true
					iterators[numInput]--
				}
			}
			continue
		}
		iterators[minI]++
		if iterators[minI] == len(input[minI]) {
			ends[minI] = true
			iterators[minI]--
		}
	}
	return res
}

package go_facet_indexer

func Invert(input []int) {
	n := len(input) / 2
	for i := 0; i < n; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
}

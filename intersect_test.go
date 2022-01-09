package go_facet_indexer

import "testing"

func TestIntersect(t *testing.T) {
	tc := [][]int{
		// {},
		{1, 2, 3, 4, 5, 6, 7},
		{-1, 2, 4, 6},
		{1, 2, 3, 4},
		{1, 1, 1, 2, 2, 2, 3, 3, 4, 4},
	}
	t.Log(Intersect(tc...))
}

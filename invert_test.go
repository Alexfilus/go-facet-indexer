package go_facet_indexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Invert(t *testing.T) {
	tcs := [][]int{{1, 2, 3}, {1, 2, 3, 4}, {1, 2, 3, 4, 5}}
	tcsRes := [][]int{{3, 2, 1}, {4, 3, 2, 1}, {5, 4, 3, 2, 1}}
	for i, tc := range tcs {
		Invert(tc)
		assert.Equal(t, tc, tcsRes[i])
	}
}

package go_facet_indexer

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	tc := []int{5, 1, 3, 7, 1, 2}
	sort.Ints(tc)
	t.Log(tc)
	t.Log(1, sort.SearchInts(tc, 1))
	t.Log(2, sort.SearchInts(tc, 2))
	t.Log(5, sort.SearchInts(tc, 5))
	t.Log(7, sort.SearchInts(tc, 7))
	t.Log(6, sort.SearchInts(tc, 6))
	t.Log(4, sort.SearchInts(tc, 4))
	t.Log(8, sort.SearchInts(tc, 8))
	t.Log(0, sort.SearchInts(tc, 0))
	t.Log(-1, sort.SearchInts(tc, -1))
}

func TestInt64(t *testing.T) {
	tc := []int64{20, 1, 15, 2, 15, 7, 8, 9, 2, 2, 22, 3}
	t.Log(tc)
	obj := Int64{}
	for i, v := range tc {
		obj.Add(v, i)
	}
	t.Log(obj)
	t.Log(1, obj.FindEqual(1))
	t.Log(2, obj.FindEqual(2))
	t.Log(3, obj.FindEqual(3))
	t.Log(7, obj.FindEqual(7))
	t.Log(8, obj.FindEqual(8))
	t.Log(9, obj.FindEqual(9))
	t.Log(15, obj.FindEqual(15))
	t.Log(20, obj.FindEqual(20))
	t.Log(22, obj.FindEqual(22))
	t.Log(-1, obj.FindEqual(-1))
	t.Log(10, obj.FindEqual(10))
	t.Log(30, obj.FindEqual(30))
}

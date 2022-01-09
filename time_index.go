package go_facet_indexer

import "time"

type Time struct {
	k []int64
	v [][]int
}

func (i *Time) Add(val time.Time, index int) {
	l := len(i.k)
	if l == 0 {
		i.k = []int64{val.UnixNano()}
		i.v = [][]int{{index}}
		return
	}
	idx, pos := i.search(val.UnixNano())
	if idx == -1 {
		if pos == 0 {
			i.k = append([]int64{val.UnixNano()}, i.k...)
			i.v = append([][]int{{index}}, i.v...)
			return
		}
		if pos == l {
			i.k = append(i.k, val.UnixNano())
			i.v = append(i.v, []int{index})
			return
		}
		i.k = append(i.k[0:pos], append([]int64{val.UnixNano()}, i.k[pos:]...)...)
		i.v = append(i.v[0:pos], append([][]int{{index}}, i.v[pos:]...)...)
		return
	}
	i.v[idx] = append(i.v[idx], index)
}

func (i *Time) Find(val time.Time) []int {
	idx, _ := i.search(val.UnixNano())
	if idx == -1 {
		return []int{}
	}
	return i.v[idx]
}

func (i *Time) search(target int64) (int, int) {
	start := 0
	end := len(i.k) - 1
	for start <= end {
		mid := (start + end) >> 1
		if i.k[mid] == target {
			return mid, -1 // found
		} else if i.k[mid] < target {
			start = mid + 1
		} else if i.k[mid] > target {
			end = mid - 1
		}
	}
	return -1, start
}

package go_facet_indexer

type Int64 struct {
	k []int64
	v [][]int
}

func (i *Int64) Add(val int64, index int) {
	l := len(i.k)
	if l == 0 {
		i.k = []int64{val}
		i.v = [][]int{{index}}
		return
	}
	idx, pos := i.search(val)
	if idx == -1 {
		if pos == 0 {
			i.k = append([]int64{val}, i.k...)
			i.v = append([][]int{{index}}, i.v...)
			return
		}
		if pos == l {
			i.k = append(i.k, val)
			i.v = append(i.v, []int{index})
			return
		}
		i.k = append(i.k[0:pos], append([]int64{val}, i.k[pos:]...)...)
		i.v = append(i.v[0:pos], append([][]int{{index}}, i.v[pos:]...)...)
		return
	}
	i.v[idx] = append(i.v[idx], index)
}

func (i *Int64) FindEqual(val int64) []int {
	idx, _ := i.search(val)
	if idx == -1 {
		return []int{}
	}
	return i.v[idx]
}

func (i *Int64) FindGT(val int64) (res []int) {
	idx, _ := i.search(val)
	if idx == -1 || idx == len(i.k)-1 {
		return []int{}
	}
	for iter := idx + 1; iter < len(i.k); iter++ {
		res = append(res, i.v[iter]...)
	}
	return res
}

func (i *Int64) FindGTE(val int64) (res []int) {
	idx, _ := i.search(val)
	if idx == -1 {
		return []int{}
	}
	for iter := idx; iter < len(i.k); iter++ {
		res = append(res, i.v[iter]...)
	}
	return res
}

func (i *Int64) FindLT(val int64) (res []int) {
	idx, _ := i.search(val)
	if idx == -1 || idx == 0 {
		return []int{}
	}
	for iter := 0; iter < idx; iter++ {
		res = append(res, i.v[iter]...)
	}
	return res
}

func (i *Int64) FindLTE(val int64) (res []int) {
	idx, _ := i.search(val)
	if idx == -1 {
		return []int{}
	}
	for iter := 0; iter <= idx; iter++ {
		res = append(res, i.v[iter]...)
	}
	return res
}

func (i *Int64) FindRange(low, high int64) (res []int) {
	if low > high {
		return
	}
	idxLow, _ := i.search(low)
	if idxLow == -1 {
		return
	}
	for iter := idxLow; i.k[iter] <= high; iter++ {
		res = append(res, i.v[iter]...)
	}
	return res
}

func (i *Int64) search(target int64) (int, int) {
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

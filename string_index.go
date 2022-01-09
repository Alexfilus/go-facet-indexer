package go_facet_indexer

type String struct {
	k []string
	v [][]int
}

func (i *String) Add(val string, index int) {
	l := len(i.k)
	if l == 0 {
		i.k = []string{val}
		i.v = [][]int{{index}}
		return
	}
	idx, pos := i.search(val)
	if idx == -1 {
		if pos == 0 {
			i.k = append([]string{val}, i.k...)
			i.v = append([][]int{{index}}, i.v...)
			return
		}
		if pos == l {
			i.k = append(i.k, val)
			i.v = append(i.v, []int{index})
			return
		}
		i.k = append(i.k[0:pos], append([]string{val}, i.k[pos:]...)...)
		i.v = append(i.v[0:pos], append([][]int{{index}}, i.v[pos:]...)...)
		return
	}
	i.v[idx] = append(i.v[idx], index)
}

func (i *String) Find(val string) []int {
	idx, _ := i.search(val)
	if idx == -1 {
		return []int{}
	}
	return i.v[idx]
}

func (i *String) search(target string) (int, int) {
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

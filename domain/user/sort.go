package user

type Sort func(i *Index) []int

func SortByIDAsc() Sort {
	return func(i *Index) []int {
		return i.idxID.SortAsc()
	}
}

func SortByIDDesc() Sort {
	return func(i *Index) []int {
		return i.idxID.SortDesc()
	}
}

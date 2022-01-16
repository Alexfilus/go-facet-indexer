package user

type Filter func(i *Index) []int

func FilterNameEq(val string) Filter {
	return func(i *Index) []int {
		return i.idxName.Find(val)
	}
}

func FilterIDGTE(val int64) Filter {
	return func(i *Index) []int {
		return i.idxID.FindGTE(val)
	}
}

func FilterIDRange(low, high int64) Filter {
	return func(i *Index) []int {
		return i.idxID.FindRange(low, high)
	}
}

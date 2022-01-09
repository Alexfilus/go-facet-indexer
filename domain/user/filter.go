package user

type Filter func(i *Index) []int

func FilterNameEq(val string) Filter {
	return func(i *Index) []int {
		return i.idxName.Find(val)
	}
}

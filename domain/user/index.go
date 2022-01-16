package user

import (
	gofacetindexer "go-facet-indexer"
)

type Index struct {
	collection   []User
	idxID        gofacetindexer.Int64
	idxName      gofacetindexer.String
	idxLastName  gofacetindexer.String
	idxBirthDate gofacetindexer.Time
	idxCreatedAt gofacetindexer.Time
}

func New() *Index {
	return &Index{
		collection:   []User{},
		idxID:        gofacetindexer.Int64{},
		idxName:      gofacetindexer.String{},
		idxLastName:  gofacetindexer.String{},
		idxBirthDate: gofacetindexer.Time{},
		idxCreatedAt: gofacetindexer.Time{},
	}
}

func (i *Index) Add(user User) {
	i.collection = append(i.collection, user)
	l := len(i.collection) - 1
	i.idxID.Add(user.ID, l)
	i.idxName.Add(user.Name, l)
	i.idxLastName.Add(user.LastName, l)
}

func (i *Index) GetByID(val int64) []User {
	return i.collect(i.idxID.FindEqual(val))
}

func (i *Index) GetByName(val string) []User {
	return i.collect(i.idxName.Find(val))
}

func (i *Index) GetByLastName(val string) []User {
	return i.collect(i.idxLastName.Find(val))
}

func (i *Index) WhereIDGTE(val int64) []User {
	return i.collect(i.idxID.FindGTE(val))
}

func (i *Index) SortByIDAsc() []User {
	return i.collect(i.idxID.SortAsc())
}

func (i *Index) SortByIDDesc() []User {
	return i.collect(i.idxID.SortDesc())
}

func (i *Index) Search() *Search {
	return &Search{
		i: i,
	}
}

func (i *Index) Len() int {
	return len(i.collection)
}

func (i *Index) collect(idxSlice []int) []User {
	res := make([]User, len(idxSlice))
	for iter, idx := range idxSlice {
		res[iter] = i.collection[idx]
	}
	return res
}

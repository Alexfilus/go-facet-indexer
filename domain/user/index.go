package user

import gofacetindexer "go-facet-indexer"

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
	idxSlice := i.idxID.FindEqual(val)
	res := make([]User, len(idxSlice))
	for iter, idx := range idxSlice {
		res[iter] = i.collection[idx]
	}
	return res
}

func (i *Index) GetByName(val string) []User {
	idxSlice := i.idxName.Find(val)
	res := make([]User, len(idxSlice))
	for iter, idx := range idxSlice {
		res[iter] = i.collection[idx]
	}
	return res
}

func (i *Index) GetByLastName(val string) []User {
	idxSlice := i.idxLastName.Find(val)
	res := make([]User, len(idxSlice))
	for iter, idx := range idxSlice {
		res[iter] = i.collection[idx]
	}
	return res
}

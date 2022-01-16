package user

import (
	gofacetindexer "go-facet-indexer"
)

type Search struct {
	i       *Index
	filters []Filter
	sort    Sort
}

func (s *Search) Do() []User {
	idsFiltered := make([][]int, len(s.filters))
	for iter, f := range s.filters {
		idsFiltered[iter] = f(s.i)
	}
	ids := gofacetindexer.Intersect(idsFiltered...)
	if s.sort != nil {
		idsSorted := s.sort(s.i)
		idsMap := make(map[int]struct{}, len(ids))
		for _, id := range ids {
			idsMap[id] = struct{}{}
		}
		res := make([]int, 0, len(ids))
		for _, id := range idsSorted {
			if _, ok := idsMap[id]; ok {
				res = append(res, id)
			}
		}
		return s.i.collect(res)
	}
	return s.i.collect(ids)
}

func (s *Search) WithFilter(filter Filter) *Search {
	s.filters = append(s.filters, filter)
	return s
}

func (s *Search) WithSort(sort Sort) *Search {
	s.sort = sort
	return s
}

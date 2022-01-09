package user

import (
	"testing"
	"time"
)

var TC = []User{
	{
		ID:        1,
		Name:      "Alex",
		LastName:  "Fil",
		BirthDate: time.Date(1991, 05, 29, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2016, 10, 4, 1, 3, 4, 0, time.UTC),
	},
	{
		ID:        2,
		Name:      "Kate",
		LastName:  "Ber",
		BirthDate: time.Date(1991, 12, 03, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2017, 3, 13, 1, 3, 4, 0, time.UTC),
	},
	{
		ID:        3,
		Name:      "Bob",
		LastName:  "Green",
		BirthDate: time.Date(1998, 1, 2, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2020, 10, 4, 1, 3, 4, 0, time.UTC),
	},
	{
		ID:        4,
		Name:      "Sam",
		LastName:  "Walker",
		BirthDate: time.Date(2005, 11, 11, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2021, 10, 4, 1, 3, 4, 0, time.UTC),
	},
	{
		ID:        5,
		Name:      "Jho",
		LastName:  "Show",
		BirthDate: time.Date(1992, 3, 4, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2017, 1, 4, 1, 3, 4, 0, time.UTC),
	},
	{
		ID:        6,
		Name:      "Alice",
		LastName:  "Bunny",
		BirthDate: time.Date(1994, 7, 6, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2018, 3, 2, 10, 3, 4, 0, time.UTC),
	},
	{
		ID:        7,
		Name:      "Tom",
		LastName:  "Lake",
		BirthDate: time.Date(1990, 11, 12, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2014, 10, 8, 1, 13, 4, 0, time.UTC),
	},
	{
		ID:        8,
		Name:      "Alex",
		LastName:  "Morgan",
		BirthDate: time.Date(1990, 3, 25, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2016, 3, 8, 1, 13, 4, 0, time.UTC),
	},
	{
		ID:        9,
		Name:      "Jimmy",
		LastName:  "Wong",
		BirthDate: time.Date(1995, 2, 8, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2015, 7, 4, 1, 3, 4, 0, time.UTC),
	},
	{
		ID:        10,
		Name:      "Tony",
		LastName:  "Stark",
		BirthDate: time.Date(1992, 1, 7, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2015, 12, 4, 1, 3, 4, 0, time.UTC),
	},
	{
		ID:        11,
		Name:      "Gary",
		LastName:  "Oldman",
		BirthDate: time.Date(1980, 1, 2, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2012, 8, 6, 5, 4, 4, 0, time.UTC),
	},
}

func TestUser(t *testing.T) {
	u := New()
	for _, tc := range TC {
		u.Add(tc)
	}
	t.Log(u.GetByID(9))
	t.Log(u.GetByName("Alex"))
	t.Log(u.GetByLastName("Oldman"))
}

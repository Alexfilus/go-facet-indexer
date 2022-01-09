package user

import "time"

//go:generate go run ../../gen.go

type User struct {
	ID        int64
	Name      string
	LastName  string
	BirthDate time.Time
	CreatedAt time.Time
}

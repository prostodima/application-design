package model

import (
	"time"
)

type (
	Hotel string
	Room  string
	User  string
)

type Order struct {
	ID    string
	Hotel Hotel
	Room  Room
	User  User
	From  time.Time
	To    time.Time
}

func (o Order) IsEmpty() bool {
	return o.ID == ""
}

type Rooms map[Hotel][]Room

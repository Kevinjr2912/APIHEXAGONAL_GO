package entities

import "time"

type Student struct {
	Id          int64
	Name        string
	Age         uint8
	PhoneNumber uint64
	UpdateAt    time.Time
}
package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category    Category
	Stock       int64
	Description string
	Createdat   time.Time
	Updatedat   time.Time
}

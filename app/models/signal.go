package models

import "time"

type Signal struct {
	Symbol    string
	Value     string
	CreatedAt time.Time `db:"created_at"`
}

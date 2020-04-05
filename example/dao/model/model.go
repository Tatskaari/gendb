package model

import "time"

type User struct {
	ID string `db:"id"`
	CreatedTime time.Time `db:"created_timestamp"`
}

type UserVersion struct {
	ID string `db:"id"`
	UserID string `db:"user_id"`
	Active bool `db:"active"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	CreateTimestamp time.Time `db:"created_timestamp"`
}


package model

import "time"

//go:generate recgen --file=model.go --record_struct=User --table_name="users"
type User struct {
	ID          string    `db:"id"`
	CreatedTime time.Time `db:"created_timestamp"`
}

//go:generate recgen --file=model.go --record_struct=UserVersion --table_name="user_versions"
type UserVersion struct {
	ID              string    `db:"id"`
	UserID          string    `db:"user_id"`
	Active          bool      `db:"active"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	CreateTimestamp time.Time `db:"created_timestamp"`
}

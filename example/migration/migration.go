package migration

import "database/sql"

const createUserTable = `
	CREATE TABLE users (
		id TEXT PRIMARY KEY,
		created_timestamp TIMESTAMP
	)
`

const createUserVersionsTable = `
	CREATE TABLE user_versions (
		id TEXT PRIMARY KEY,
		created_timestamp TIMESTAMP,
		active INTEGER,
		user_id string NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
`

func Migrate(db *sql.DB) error {
	_, err := db.Exec(createUserTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createUserVersionsTable)
	if err != nil {
		return err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	return err
}
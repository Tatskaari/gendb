package model

import (
	"github.com/tatskaari/gendb/example/dao/model"
	"time"
)

type User struct {
	ID string
	CreatedTime time.Time
	UpdatedTime time.Time
	FirstName string
	LastName string
}

func DatabaseUserToUser(user *model.User, version *model.UserVersion) *User {
	return &User{
		ID: user.ID,
		CreatedTime: user.CreatedTime,
		UpdatedTime: version.CreateTimestamp,
		FirstName: version.FirstName,
		LastName: version.LastName,
	}
}
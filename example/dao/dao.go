package dao

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tatskaari/gendb/builder"
	"github.com/tatskaari/gendb/example/dao/model"
	"github.com/tatskaari/gendb/executor"
	"time"
)

const (
	userTable = "users"
	userVersionsTable = "user_versions"
)

type DAO struct {
	ex *executor.Executor
}

func New(ex *executor.Executor) *DAO {
	return &DAO{ex:ex}
}

func (dao *DAO) AddUser(user *model.User) error {
	_, err := dao.ex.Insert().Into(userTable).Values(map[string]interface{}{
		"id": builder.Bind(user.ID),
		"created_timestamp": time.Now(),
	}).Exec()
	if err != nil {
		return fmt.Errorf("failed to insert user %v: %w", user, err)
	}
	return nil
}

func (dao *DAO) AddUserVersion(version *model.UserVersion) error {
	_, err := dao.ex.Update(userVersionsTable).
		Set("active", false).
		Where(builder.Eq("user_id", builder.Bind(version.UserID))).Exec()
	if err != nil {
		return err
	}

	_, err = dao.ex.Insert().Into(userVersionsTable).Values(map[string]interface{}{
		"id": builder.Bind(uuid.New().String()),
		"created_timestamp": time.Now(),
		"user_id": builder.Bind(version.UserID),
		"active": true,
		"first_name": builder.Bind(version.FirstName),
		"last_name": builder.Bind(version.LastName),
	}).Exec()
	return err
}

func (dao *DAO) GetUser(id string) (*model.User, error) {
	var user model.User
	err := dao.ex.Select("*").
		From(userTable).
		Where(builder.Eq("id", builder.Bind(id))).
		Get(&user)
	return &user, err
}

func (dao *DAO) GetActiveUserVersion(id string) (*model.UserVersion, error) {
	var version model.UserVersion
	err := dao.ex.Select("*").
		From(userVersionsTable).
		Where(builder.Eq("user_id", builder.Bind(id))).
		And("active").
		Get(&version)
	return &version, err
}
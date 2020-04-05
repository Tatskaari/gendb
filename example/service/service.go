package service

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tatskaari/gendb/example/dao"
	daomodel "github.com/tatskaari/gendb/example/dao/model"
	"github.com/tatskaari/gendb/example/service/model"
)

type Service struct {
	dbCon *sqlx.DB
}

func New(con *sqlx.DB) *Service {
	return &Service{dbCon: con}
}

func (s *Service) AddUser(user *model.User) error {
	user.ID = uuid.New().String()
	return dao.InTransaction(s.dbCon, func(dao *dao.DAO) error {
		err := dao.AddUser(&daomodel.User{ID: user.ID})
		if err != nil {
			return err
		}
		return dao.AddUserVersion(&daomodel.UserVersion{
			UserID:    user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	})
}

func (s *Service) UpdateUser(user *model.User) error {
	return dao.InTransaction(s.dbCon, func(dao *dao.DAO) error {
		return dao.AddUserVersion(&daomodel.UserVersion{
			UserID:    user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	})
}

func (s *Service) GetUser(userID string) (*model.User, error) {
	var user *daomodel.User
	var version *daomodel.UserVersion

	err := dao.InTransaction(s.dbCon, func(dao *dao.DAO) error {
		var err error
		user, err = dao.GetUser(userID)
		if err != nil {
			return err
		}
		version, err = dao.GetActiveUserVersion(userID)
		return err
	})

	if err != nil {
		return nil, err
	}
	return model.DatabaseUserToUser(user, version), nil
}

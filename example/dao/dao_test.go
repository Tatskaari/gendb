package dao_test

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/tatskaari/gendb/example/dao"
	"github.com/tatskaari/gendb/example/dao/model"
	"github.com/tatskaari/gendb/example/migration"
	servicemodel "github.com/tatskaari/gendb/example/service/model"
	"github.com/tatskaari/gendb/executor"
	"github.com/tatskaari/gendb/sqlizer"
	"testing"
)

type daoSuite struct {
	suite.Suite
	dao *dao.DAO
}

func TestDAO(t *testing.T) {
	suite.Run(t, new(daoSuite))
}

func (s *daoSuite) SetupTest() {
	con, err := sqlx.Open("sqlite3", ":memory:")
	s.NoError(err, "failed to start database")

	err = migration.Migrate(con.DB)
	s.NoError(err)

	s.dao = dao.New(executor.New(con, new(sqlizer.StandardSqlizer)))
}

func (s *daoSuite) TestCreateUser() {
	expectedUser := &servicemodel.User{
		ID:        uuid.New().String(),
		FirstName: "John",
		LastName:  "Doe",
	}

	err := s.dao.AddUser(&model.User{ID: expectedUser.ID})
	s.Require().NoError(err)

	user, err := s.dao.GetUser(expectedUser.ID)
	s.Require().NoError(err)

	s.Equal(expectedUser.ID, user.ID)
	s.NotZero(user.CreatedTime)

	err = s.dao.AddUserVersion(&model.UserVersion{
		UserID:    expectedUser.ID,
		Active:    true,
		FirstName: expectedUser.FirstName,
		LastName:  expectedUser.LastName,
	})
	s.Require().NoError(err)

	version, err := s.dao.GetActiveUserVersion(expectedUser.ID)
	s.Require().NoError(err)

	s.Equal(expectedUser.ID, version.UserID)
	s.NotZero(version.CreateTimestamp)
	s.NotZero(version.ID)
	s.Equal(expectedUser.FirstName, version.FirstName)
	s.Equal(expectedUser.LastName, version.LastName)

	updatedUser := &model.UserVersion{
		UserID:        expectedUser.ID,
		FirstName: "Jane",
		LastName:  "Smith",
	}
	err = s.dao.AddUserVersion(updatedUser)
	s.Require().NoError(err)

	version, err = s.dao.GetActiveUserVersion(user.ID)
	s.Require().NoError(err)

	s.Equal(updatedUser.FirstName, version.FirstName)
	s.Equal(updatedUser.LastName, version.LastName)
}

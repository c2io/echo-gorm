package model

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	mock sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	var (
		db  *sql.DB
		err error
	)

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db, SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	user := &User{Username: "c", Password: "io"}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`username`,`password`) VALUES (?,?)").
		WithArgs(user.Username, user.Password).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	DB.Create(user)
}

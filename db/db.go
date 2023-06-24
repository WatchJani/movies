package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//mysql://be5cbbfc5d050e:d4d84566@eu-cdbr-west-03.cleardb.net/heroku_d0920392b7eace1?reconnect=true

// user : be5cbbfc5d050e
// password : d4d84566
// hostName : eu-cdbr-west-03.cleardb.net
// data base: heroku_d0920392b7eace1

func Init() error {
	var err error
	db, err = sqlx.Open("mysql", "be5cbbfc5d050e:d4d84566@tcp(eu-cdbr-west-03.cleardb.net)/heroku_d0920392b7eace1")
	if err != nil {
		return err
	}

	return nil
}

func Select(dest interface{}, query string, args ...interface{}) error {
	return db.Select(dest, query, args...)
}

func Insert(query string, args ...any) (sql.Result, error) {
	return db.Exec(query, args...)
}

func Delete(query string, args ...any) (sql.Result, error) {
	return db.Exec(query, args...)
}

func Update(query string, args ...any) (sql.Result, error) {
	return db.Exec(query, args...)
}

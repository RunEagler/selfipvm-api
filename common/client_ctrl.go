package common

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/lib/pq"
)

//ConnectPostgres :connect postgres
func ConnectPostgres(user string, dbName string) (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable",
		user,
		dbName,
	))
	if err != nil {
		return nil, err
	}
	return db, nil
}

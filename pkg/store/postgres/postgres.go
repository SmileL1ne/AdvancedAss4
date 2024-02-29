package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func OpenDB(dbInfo *Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	"runtime"
)

func GetDbConnection() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s dbname=%s host=%s port=%d password=%s sslmode=disable",
			Config.GetString("database.username"),
			Config.GetString("database.database"),
			Config.GetString("database.host", "localhost"),
			Config.GetInt64("database.port", 5432),
			Config.GetString("database_password"),
		),
	)

	// coreCount * 2 is apparently a good number of connections according to:
	// http://wiki.postgresql.org/wiki/Number_Of_Database_Connections
	//      #How_to_Find_the_Optimal_Database_Connection_Pool_Size
	db.SetMaxIdleConns(runtime.NumCPU() * 2)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("helpers/db.go: Database connection failed: %v", err.Error()))
	}

	return db, err
}

func GetDbTransaction() (*sql.Tx, error) {
	db, err := GetDbConnection()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("helpers/db.go: Could not start a transaction: %v", err.Error()))
	}

	return tx, err
}

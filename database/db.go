package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var AssetDB *sqlx.DB

func ConnectAndMigrate(host, port, user, password, dbName, sslmode string) error {

	cntstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslmode)

	db, dbErr := sqlx.Open("postgres", cntstr)

	if dbErr != nil {
		return dbErr
	}

	if pingErr := db.Ping(); pingErr != nil {
		return pingErr
	}

	AssetDB = db
	return migrateUp(db)
}

func migrateUp(db *sqlx.DB) error {

	driver, driverErr := postgres.WithInstance(db.DB, &postgres.Config{})
	if driverErr != nil {
		return driverErr
	}

	m, migrateErr := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres",
		driver,
	)

	if migrateErr != nil {
		return migrateErr
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func Tx(fn func(tx *sqlx.Tx) error) error {

	tx, err := AssetDB.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	defer func() {
		if err != nil {
			if Txerr := tx.Rollback(); Txerr != nil {
				logrus.Errorf("failed to rollback transation: %v", Txerr)
			}
			return
		}
		if Txerr := tx.Commit(); Txerr != nil {
			logrus.Errorf("failed to commit transation: %v", Txerr)
		}
	}()

	err = fn(tx)

	return err
}

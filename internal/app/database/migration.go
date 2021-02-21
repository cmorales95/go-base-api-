package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/labstack/gommon/log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	label = "Run Migrations"
)

// ConfigureDatabaseSQL connects to db and make the migrations
func ConfigureDatabaseSQL() {
	runMigration()
}

func runMigration() {
	dbClient := Client

	driver, err := postgres.WithInstance(dbClient, &postgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("%s | Failed creating database: %s", label, err))
	}

	migration, err := migrate.NewWithDatabaseInstance("file://internal/app/database/migration", "postgres", driver)
	if err != nil {
		panic(fmt.Sprintf("%s | Failed creating migrations instance: %s", label, err))
	}

	if err = migration.Up(); err != nil {
		log.Errorf("%s | Failed running migrations: %v", label, err)
	}
	log.Info("%s | Database migrations has been succesfully")

}

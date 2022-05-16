package main_test

//go:generate sqlboiler --wipe psql

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func ExampleQueries() {
	psql, err := sql.Open("pgx", "host=postgis user=root dbname=experiment password=hunter2 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	runMigrations(psql)

	// boil.SetDB(psql)

	// output:
}

func runMigrations(psql *sql.DB) {
	driver, err := postgres.WithInstance(psql, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed getting postgres instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("failed initializing migrations: %v", err)
	}

	err = m.Down()
	if err != nil {
		log.Printf("failed dropping database: %v", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("failed running migrations: %v", err)
	}
}

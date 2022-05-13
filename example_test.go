package main_test

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func ExampleQueries() {
	// ctx := context.Background()

	psql, err := sql.Open("pgx", "host=postgis user=root dbname=experiment password=hunter2 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	runMigrations(psql)
}

func runMigrations(psql *sql.DB) {
	driver, err := postgres.WithInstance(psql, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed getting postgres instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("failed initializing migrations: %v", err)
	}

	err = m.Down()
	if err != nil {
		log.Fatalf("failed dropping database: %v", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("failed running migrations: %v", err)
	}
}

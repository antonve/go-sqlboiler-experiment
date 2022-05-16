package main_test

//go:generate sqlboiler --wipe psql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/antonve/go-sqlboiler-experiment/models"
	"github.com/paulmach/orb"
	"github.com/volatiletech/sqlboiler/v4/boil"

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

	ctx := context.Background()

	// create restaurants
	cocoId := createRestaurant(psql, "CoCo Ichibanya Ebisu", 35.64699825984844, 139.71194575396922)
	shakeId := createRestaurant(psql, "SHAKE SHACK Ebisu", 35.64669248211187, 139.70949784477963)
	ichiranId := createRestaurant(psql, "Ichiran Ramen Shinjuku", 35.69079988476277, 139.70286473414785)
	torikiId := createRestaurant(psql, "Torikizoku Shinjuku", 35.68918337273537, 139.70249991934935)

	restaurants, err := models.Restaurants().All(ctx, psql)
	fmt.Print(restaurants, cocoId, shakeId, ichiranId, torikiId)

	// output:
	// []
}

func createRestaurant(psql *sql.DB, name string, long, lat float64) int64 {
	var r models.Restaurant
	r.Name = name
	r.Location = orb.Point{long, lat}

	err := r.Insert(context.Background(), psql, boil.Infer())
	if err != nil {
		log.Fatalf("cannot create restaurant: %v", err)
	}

	return r.ID
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

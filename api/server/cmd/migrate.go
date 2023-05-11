package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/urfave/cli/v2"
)

var (

	// Number of migrations to run, it should either be empty
	// or a string such as "+n" or "-n" where n is an integer
	migrationsNum int

	// The path from which to pick up the migrations files
	migrationsPath string

	migrateCommand = &cli.Command{
		Name:   "migrate",
		Usage:  "Run migrations",
		Action: migrateCmd,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "num",
				Usage: "number of migrations to apply of the format +n" +
					"or -n. Where n is an integer. Without a value all" +
					"migrations will be applied.",
				Required:    false,
				Destination: &migrationsNum,
			},
			&cli.StringFlag{
				Name: "migrationsPath",
				Usage: "number of migrations to apply of the format +n" +
					"or -n. Where n is an integer. Without a value all" +
					"migrations will be applied.",
				Required:    true,
				Destination: &migrationsPath,
			},
		},
	}
)

func migrateCmd(c *cli.Context) error {
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		))
	if err != nil {
		log.Fatalf("failed to connect db:%s", err.Error())
	}

	fmt.Println(fmt.Sprintf("%s   %s   %s  %s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
	))

	path := fmt.Sprintf("file://%s", migrationsPath)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed driver connect: %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		path,
		"postgres",
		driver)
	if err != nil {
		log.Fatalf("failed db connect: %s", err.Error())
	}

	if migrationsNum == 0 {
		if err := m.Up(); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("no change")
			} else {
				log.Fatalf("failed migration: %s", err.Error())
			}
		}
	} else {
		if err := m.Steps(migrationsNum); err != nil {
			log.Fatalf("failed migration steps: %s", err.Error())
		}
	}

	return nil
}

func init() {
	register(migrateCommand)
}

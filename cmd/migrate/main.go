package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/aji/chi-go-template/internal/config"
	"github.com/aji/chi-go-template/migrations"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg, err := config.Load(".env")
	if err != nil {
		cfg, err = config.Load("")
		if err != nil {
			fmt.Println("config error:", err)
			os.Exit(1)
		}
	}

	src, err := iofs.New(migrations.FS, ".")
	if err != nil {
		fmt.Println("source error:", err)
		os.Exit(1)
	}

	var db *sql.DB
	switch cfg.Database.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
		db, err = sql.Open("mysql", dsn)
	case "postgres":
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
		db, err = sql.Open("postgres", dsn)
	case "sqlite":
		db, err = sql.Open("sqlite3", cfg.Database.Name)
	default:
		fmt.Println("unknown db driver")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("db open error:", err)
		os.Exit(1)
	}
	defer db.Close()

	var driver database.Driver
	switch cfg.Database.Driver {
	case "mysql":
		d, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			fmt.Println("driver error:", err)
			os.Exit(1)
		}
		driver = d
	case "postgres":
		d, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			fmt.Println("driver error:", err)
			os.Exit(1)
		}
		driver = d
	case "sqlite":
		d, err := sqlite3.WithInstance(db, &sqlite3.Config{})
		if err != nil {
			fmt.Println("driver error:", err)
			os.Exit(1)
		}
		driver = d
	}

	m, err := migrate.NewWithInstance("iofs", src, cfg.Database.Driver, driver)
	if err != nil {
		fmt.Println("migrate init error:", err)
		os.Exit(1)
	}

	cmd := "up"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("migrate up error:", err)
			os.Exit(1)
		}
		fmt.Println("migrations applied")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("migrate down error:", err)
			os.Exit(1)
		}
		fmt.Println("migrations rolled back")
	default:
		fmt.Println("usage: migrate [up|down]")
		os.Exit(1)
	}
}

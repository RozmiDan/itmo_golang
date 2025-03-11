package main

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	cfgURL := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable&pool_max_conns=10"

	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, cfgURL)

	if err != nil {
		log.Fatalln(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("cant set dialect", err)
	}

	db := stdlib.OpenDBFromPool(connPool)

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("cant up migration", err)
	}

	fmt.Println("Migrations applied successfully!")
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	cfgURL := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable&pool_max_conns=10"

	dbPool, err := pgxpool.New(ctx, cfgURL)

	if err != nil {
		log.Fatal("can't create psgr connection")
	}

	insertUsers := `INSERT INTO users (first_name, second_name) 
						VALUES('Daniel', 'Romirski'),
						('Alex', 'Chend'),
						('George', 'Makle')
						RETURNING id
						;`

	row := dbPool.QueryRow(ctx, insertUsers)

	var i int
	row.Scan(&i)
	fmt.Println(i)

	defer dbPool.Close()
}

// func main() {
// 	ctx := context.Background()

// 	cfgURL := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"

// 	conn, err := pgx.Connect(ctx, cfgURL)

// 	if err != nil {
// 		log.Fatal("can't create psgr connection")
// 	}

// 	createUsers := `CREATE TABLE IF NOT EXISTS users(
// 						id bigserial primary key,
// 						first_name varchar(50) not null,
// 						second_name varchar(50) not null
// 					);`

// 	_, err = conn.Exec(ctx, createUsers)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer conn.Close(ctx)
// }

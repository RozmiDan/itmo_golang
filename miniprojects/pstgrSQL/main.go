package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// DSN (Data Source Name) в формате:
	// "postgres://username:password@host:port/dbname?sslmode=disable"
	connStr := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка открытия подключения: %v", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе: %v", err)
	}
	fmt.Println("Успешное подключение к базе!")

	// Пример выполнения запроса
	{
		var now string
		err = db.QueryRow("SELECT NOW()").Scan(&now)
		if err != nil {
			log.Fatalf("Ошибка запроса: %v", err)
		}
		fmt.Printf("Сервер базы данных сообщает, что сейчас: %s\n", now)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(2) * time.Second)

	{
		var query = `CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL, 
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`
		
		res, err := db.ExecContext(ctx, query)
		if err != nil{
			log.Fatalln("Error while trying create table")
		}
		fmt.Println(res.RowsAffected())
	}

	{
		insertQuery := `
		INSERT INTO users(name, email, created_at)
		VALUES($1, $2, now())
		RETURNING id
		`
		res, err := db.ExecContext(ctx, insertQuery, "Said", "lffffs@mail.com")
		if err != nil{
			log.Fatalln("Error while trying create table")
		}
		fmt.Println(res.RowsAffected())
	}

	{
		selectQuery := `
			SELECT email FROM users WHERE id = $1
		`
		var result string
		row := db.QueryRowContext(ctx, selectQuery, 2)
		err := row.Scan(&result)
		if err == sql.ErrNoRows{
			fmt.Println("row is not found")
		} else if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("mail: ", result)
	}

	{
		selectQuery := `
			SELECT email FROM users WHERE id > 0
		`
		rows, err := db.QueryContext(ctx, selectQuery)
		if err != nil{
			log.Fatalln("Cant do query")
		}
		defer rows.Close()
		
		mails := make([]string, 0)

		for rows.Next() {
			var mail string
			err = rows.Scan(&mail)
			if err != nil{
				log.Fatalln("Cant read row in var")
			}
			mails = append(mails, mail)
		}
		for _, i := range(mails){
			fmt.Println(i)
		}
		
	}
	
}

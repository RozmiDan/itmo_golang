package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	cfgURL := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable&pool_max_conns=10"

	rep, err := NewRepo(cfgURL)

	if err != nil {
		log.Fatal(err)
	}

	book := Book{
		ID:                1,
		AuthorID:          2,
		Title:             "Отцы и дети",
		YearOfPublication: 1854,
	}

	book, err = rep.CreateBook(ctx, book)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(book.CreatedAt, book.UpdatedAt)
	fmt.Println("All correct")
}

type Repository struct {
	connPool *pgxpool.Pool
}

func NewRepo(url string) (*Repository, error) {
	cPool, err := pgxpool.New(context.Background(), url)

	if err != nil {
		return &Repository{}, err
	}

	return &Repository{connPool: cPool}, nil
}

func (r *Repository) CreateBook(ctx context.Context, book Book) (Book, error) {
	tx, err := r.connPool.Begin(ctx)

	if err != nil {
		return Book{}, err
	}

	defer tx.Rollback(ctx)

	const queryBook = `INSERT INTO books(title, year)
					  VALUES($1, $2)
					  RETURNING id, created_at, updated_at;`

	const queryAuthorBooks = `INSERT INTO author_books(author_id, book_id)
							 VALUES($1, $2);
							 `

	err = tx.QueryRow(ctx, queryBook, book.Title,
		book.YearOfPublication).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		return Book{}, err
	}

	_, err = tx.Exec(ctx, queryAuthorBooks, book.AuthorID, book.ID)

	if err != nil {
		return Book{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		log.Printf("Error in transaction")
		return Book{}, err
	}

	return book, nil
}

type Book struct {
	ID                int
	AuthorID          int
	Title             string
	YearOfPublication uint32
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

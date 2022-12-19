package repositories

import (
	"database/sql"
	"fmt"

	"github.com/AdrianoSantana1/marcador-pagina/backend/cmd/data"
)

type Book struct {
	db *sql.DB
}

func BookRepository(db *sql.DB) *Book {
	return &Book{db}
}

func (repository *Book) Create(bookModel *data.Book) error {
	fmt.Println("Inserindo usuário!")

	statement, err := repository.db.Prepare("INSERT INTO public.books (title, last_date_read, chapter, page) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(
		bookModel.Title,
		bookModel.Last_date_read,
		bookModel.Chapter,
		bookModel.Page,
	)

	if err != nil {
		return err
	}

	return nil
}

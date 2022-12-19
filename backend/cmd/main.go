package main

import (
	"fmt"
	"log"

	"github.com/AdrianoSantana1/marcador-pagina/backend/cmd/data"
	"github.com/AdrianoSantana1/marcador-pagina/backend/cmd/repositories"
)

func main() {
	// connect to database
	db, err := data.Connect()
	if err != nil {
		log.Fatal(err)
	}

	bookRepository := repositories.BookRepository(db)

	book := data.Book{
		Title:          "Gastaria tudo com pizza!",
		ID:             0,
		Last_date_read: "2022-12-10",
		Chapter:        12,
		Page:           123,
	}

	defer db.Close()
	if err := bookRepository.Create(&book); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserido marcador!")
}

package repository_test

import (
	"context"
	"testing"

	"main.go/pkg/models"
	"main.go/pkg/repository"
)

const connString = "postgres://limbo:10007@localhost:5432/books"

func TestCRUDBooks(t *testing.T) {
	ctx := context.Background()

	database, err := repository.New(context.Background(), connString)
	if err != nil {
		t.Fatal(err)
	}

	books, err := database.ReadBooks(ctx, 1, 100)
	length := len(books)

	b := models.Book{Name: "Бесы", Price: 15000, GenreId: 1, AuthorId: 2}
	err = database.AddBook(ctx, b)
	if err != nil {
		t.Fatal(err)
	}
	books, err = database.ReadBooks(ctx, 1, 100)
	if err != nil || length == len(books) {
		t.Fatal(err)
	}

	b = books[0]
	books[0].Price++
	err = database.UpdateBooksAt(ctx, books[0].Id, books[0])
	books, err = database.ReadBooks(ctx, 1, 100)
	if err != nil || b.Price == books[0].Price {
		t.Fatal(err)
	}
}

func TestCRUDGenres(t *testing.T) {
	ctx := context.Background()
	database, err := repository.New(ctx, connString)
	if err != nil {
		t.Fatal(err)
	}

	genres, err := database.ReadGenres(ctx, 1, 100)
	if err != nil {
		t.Fatal(err)
	}
	length := len(genres)
	g := models.Genre{Name: "Ода"}
	err = database.AddGenre(ctx, g)
	if err != nil {
		t.Fatal(err)
	}
	genres, err = database.ReadGenres(ctx, 1, 100)
	if err != nil || length == len(genres) {
		t.Fatal(err)
	}
	g = genres[0]
	g.Name = "Басня"
	err = database.UpdateGenresAt(ctx, g.Id, g)
	if err != nil {
		t.Fatal(err)
	}
	genres, err = database.ReadGenres(ctx, 1, 100)
	if err != nil || g.Name == genres[0].Name {
		t.Fatal()
	}
}

func TestCRUDAuthors(t *testing.T) {
	ctx := context.Background()
	database, err := repository.New(ctx, connString)
	if err != nil {
		t.Fatal(err)
	}

	authors, err := database.ReadAuthors(ctx, 1, 100)
	if err != nil {
		t.Fatal(err)
	}
	length := len(authors)
	a := models.Author{Name: "Г. Лавкрафт"}
	err = database.AddAuthor(ctx, a)
	if err != nil {
		t.Fatal(err)
	}
	authors, err = database.ReadAuthors(ctx, 1, 100)
	if err != nil || length == len(authors) {
		t.Fatal(err)
	}
	a = authors[0]
	a.Name = "А.П. Чехов"
	err = database.UpdateAuthorsAt(ctx, a.Id, a)
	if err != nil {
		t.Fatal(err)
	}
	authors, err = database.ReadAuthors(ctx, 1, 100)
	if err != nil || a.Name == authors[0].Name {
		t.Fatal()
	}
}

//Все тесты прошли :)))

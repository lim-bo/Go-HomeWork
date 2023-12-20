package repository

import (
	"context"

	"main.go/pkg/models"
)

func (repo *PgRepo) AddBook(ctx context.Context, book models.Book) error {
	repo.M.Lock()
	defer repo.M.Unlock()
	request := `INSERT INTO books (name, price, genre_id, author_id) values
				($1, $2, $3, $4)`
	_, err := repo.Pool.Exec(
		ctx, request,
		book.Name,
		book.Price,
		book.GenreId,
		book.AuthorId,
	)
	return err
}

func (repo *PgRepo) RemoveBook(ctx context.Context, id int) error {
	repo.M.Lock()
	defer repo.M.Unlock()

	request := `DELETE FROM books WHERE id=$1;`
	_, err := repo.Pool.Exec(ctx, request, id)
	return err
}

func (repo *PgRepo) UpdateBooksAt(ctx context.Context, id int, b models.Book) error {
	repo.M.Lock()
	defer repo.M.Unlock()

	request := `UPDATE books SET name=$1, price=$2, author_id=$3, genre_id=$4 
				WHERE id = $5;`
	_, err := repo.Pool.Exec(
		ctx, request,
		b.Name,
		b.Price,
		b.AuthorId,
		b.GenreId,
		id,
	)
	return err
}

func (repo *PgRepo) ReadBook(ctx context.Context, id int) (models.Book, error) {

	request := `SELECT id, name, price, genre_id, author_id FROM books
				WHERE id=$1;`
	b := models.Book{}
	err := repo.Pool.QueryRow(ctx, request, id).Scan(&b.Id, &b.Name, &b.Price, &b.GenreId, &b.AuthorId)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (repo *PgRepo) ReadBooks(ctx context.Context, from int, cnt int) ([]models.Book, error) {
	request := `SELECT id, name, price, genre_id, author_id FROM books
				WHERE id>=$1 LIMIT $2`
	var books []models.Book

	rows, err := repo.Pool.Query(ctx, request, from, cnt)
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var b models.Book
		err = rows.Scan(&b.Id, &b.Name, &b.Price, &b.GenreId, &b.AuthorId)
		if err != nil {
			return books, err
		}
		books = append(books, b)
	}

	return books, nil
}

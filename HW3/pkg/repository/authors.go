package repository

import (
	"context"

	"main.go/pkg/models"
)

func (repo *PgRepo) AddAuthor(ctx context.Context, a models.Author) error {
	repo.M.Lock()
	defer repo.M.Unlock()

	request := `INSERT INTO authors (name) values ($1);`
	_, err := repo.Pool.Exec(ctx, request, a.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PgRepo) RemoveAuthor(ctx context.Context, id int) error {
	repo.M.Lock()
	defer repo.M.Unlock()

	request := `DELETE FROM authors WHERE id=$1;`
	_, err := repo.Pool.Exec(ctx, request, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PgRepo) UpdateAuthorsAt(ctx context.Context, id int, g models.Author) error {
	repo.M.Lock()
	defer repo.M.Unlock()

	request := `UPDATE authors SET name=$1 WHERE id=$2;`
	_, err := repo.Pool.Exec(ctx, request, g.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PgRepo) ReadAuthor(ctx context.Context, id int) (models.Author, error) {
	a := models.Author{}
	request := `SELECT id, name FROM authors WHERE id=$1;`
	err := repo.Pool.QueryRow(ctx, request, id).Scan(&a.Id, &a.Name)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (repo *PgRepo) ReadAuthors(ctx context.Context, from int, cnt int) ([]models.Author, error) {
	request := `SELECT id, name FROM authors
				WHERE id>=$1 LIMIT $2`
	authors := []models.Author{}

	rows, err := repo.Pool.Query(ctx, request, from, cnt)
	if err != nil {
		return authors, err
	}

	for rows.Next() {
		var a models.Author
		err = rows.Scan(&a.Id, &a.Name)
		if err != nil {
			return authors, err
		}
		authors = append(authors, a)
	}

	return authors, nil
}

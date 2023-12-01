package repository

import (
	"context"

	"main.go/pkg/models"
)

func (repo *PgRepo) AddGenre(ctx context.Context, g models.Genre) error {
	repo.M.Lock()
	defer repo.M.Unlock()

	request := `INSERT INTO genres (name) values ($1);`
	_, err := repo.Pool.Exec(ctx, request, g.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PgRepo) RemoveGenre(ctx context.Context, id int) error {
	repo.M.Lock()
	defer repo.M.Unlock()

	request := `DELETE FROM genres WHERE id=$1;`
	_, err := repo.Pool.Exec(ctx, request, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PgRepo) UpdateGenresAt(ctx context.Context, id int, g models.Genre) error {
	repo.M.Lock()
	defer repo.M.Unlock()

	request := `UPDATE genres SET name=$1 WHERE id=$2;`
	_, err := repo.Pool.Exec(ctx, request, g.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PgRepo) ReadGenre(ctx context.Context, id int) (models.Genre, error) {
	g := models.Genre{}
	request := `SELECT id, name FROM genres WHERE id=$1;`
	err := repo.Pool.QueryRow(ctx, request, id).Scan(&g.Id, &g.Name)
	if err != nil {
		return g, err
	}
	return g, nil
}

func (repo *PgRepo) ReadGenres(ctx context.Context, from int, cnt int) ([]models.Genre, error) {
	request := `SELECT id, name FROM genres
				WHERE id>=$1 LIMIT $2`
	genres := []models.Genre{}

	rows, err := repo.Pool.Query(ctx, request, from, cnt)
	if err != nil {
		return genres, err
	}

	for rows.Next() {
		var g models.Genre
		err = rows.Scan(&g.Id, &g.Name)
		if err != nil {
			return genres, err
		}
		genres = append(genres, g)
	}

	return genres, nil
}

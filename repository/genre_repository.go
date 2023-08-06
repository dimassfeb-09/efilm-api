package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
)

type GenreRepository interface {
	Save(ctx context.Context, tx *sql.Tx, genre *domain.Genre) error
	Update(ctx context.Context, tx *sql.Tx, genre *domain.Genre) error
	Delete(ctx context.Context, tx *sql.Tx, ID int) error
	FindAll(ctx context.Context, db *sql.DB) ([]*domain.Genre, error)
	FindByName(ctx context.Context, db *sql.DB, name string) (*domain.Genre, error)
	FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Genre, error)
}

type GenreRepositoryaImpl struct {
}

func NewGenreRepository() GenreRepository {
	return &GenreRepositoryaImpl{}
}

func (repository *GenreRepositoryaImpl) Save(ctx context.Context, tx *sql.Tx, genre *domain.Genre) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO genres (name) VALUES ($1)", genre.Name)
	if err != nil {
		return err
	}

	return nil
}

func (repository *GenreRepositoryaImpl) Update(ctx context.Context, tx *sql.Tx, genre *domain.Genre) error {
	_, err := tx.ExecContext(ctx, "UPDATE genres SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", genre.Name, genre.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *GenreRepositoryaImpl) Delete(ctx context.Context, tx *sql.Tx, ID int) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM genres WHERE id = $1", ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New(fmt.Sprintf("Data with ID %d not found", ID))
		}
		return err
	}

	return nil
}

func (repository *GenreRepositoryaImpl) FindAll(ctx context.Context, db *sql.DB) ([]*domain.Genre, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM genres")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed get data from database")
		}
		return nil, err
	}

	var genres []*domain.Genre
	for rows.Next() {
		var genre domain.Genre
		rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt)
		genres = append(genres, &genre)
	}

	return genres, nil
}

func (repository *GenreRepositoryaImpl) FindByName(ctx context.Context, db *sql.DB, name string) (*domain.Genre, error) {
	var genre domain.Genre
	err := db.QueryRowContext(ctx, "SELECT * FROM genres WHERE name = $1", name).Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("genre with name %s not found", name)
		}
		return nil, err
	}
	return &genre, nil
}
func (repository *GenreRepositoryaImpl) FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Genre, error) {
	var genre domain.Genre
	err := db.QueryRowContext(ctx, "SELECT * FROM genres WHERE id = $1", ID).Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("genre with id %d not found", ID)
		}
		return nil, err
	}

	return &genre, nil
}

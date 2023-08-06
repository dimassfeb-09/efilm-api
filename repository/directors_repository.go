package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
)

type DirectorRepository interface {
	Save(ctx context.Context, tx *sql.Tx, director *domain.Director) error
	Update(ctx context.Context, tx *sql.Tx, director *domain.Director) error
	Delete(ctx context.Context, tx *sql.Tx, ID int) error
	FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Director, error)
	FindByName(ctx context.Context, db *sql.DB, name string) (*domain.Director, error)
	FindByNational(ctx context.Context, db *sql.DB, nationalityID int) ([]*domain.Director, error)
	FindAll(ctx context.Context, db *sql.DB) ([]*domain.Director, error)
}

type DirectorRepositoryImpl struct {
}

func NewDirectorRepository() DirectorRepository {
	return &DirectorRepositoryImpl{}
}

func (a *DirectorRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, director *domain.Director) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO directors (name, date_of_birth, nationality_id) VALUES ($1, $2, $3)", director.Name, director.DateOfBirth, director.NationalityID)
	if err != nil {
		return err
	}

	return nil
}

func (a *DirectorRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, director *domain.Director) error {
	query := "UPDATE directors SET id = $1, name = $2, date_of_birth = $3, nationality_id = $4, updated_at = $5 WHERE id = $6"
	_, err := tx.ExecContext(ctx, query, director.ID, director.Name, director.DateOfBirth, director.NationalityID, director.UpdatedAt, director.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *DirectorRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, ID int) error {
	_, err := tx.Exec("DELETE FROM directors WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *DirectorRepositoryImpl) FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Director, error) {
	var director domain.Director
	err := db.QueryRow("SELECT * FROM directors WHERE id = $1", ID).Scan(&director.ID, &director.Name, &director.DateOfBirth, &director.NationalityID, &director.CreatedAt, &director.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("sorry, id not found")
		}
		return nil, err
	}

	return &director, nil
}

func (a *DirectorRepositoryImpl) FindByName(ctx context.Context, db *sql.DB, name string) (*domain.Director, error) {
	var director domain.Director
	err := db.QueryRow("SELECT * FROM directors WHERE name = $1", name).Scan(&director.ID, &director.Name, &director.DateOfBirth, &director.NationalityID, &director.CreatedAt, &director.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("director with name %s not found", name)
		}
		return nil, err
	}

	return &director, nil
}

func (a *DirectorRepositoryImpl) FindByNational(ctx context.Context, db *sql.DB, nationalityID int) ([]*domain.Director, error) {
	rows, err := db.Query("SELECT * FROM directors WHERE nationality_id = $1", nationalityID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("nationality with ID %d not found", nationalityID)
		}
		return nil, err
	}

	var directors []*domain.Director
	for rows.Next() {
		var director domain.Director
		rows.Scan(&director.ID, &director.Name, &director.DateOfBirth, &director.NationalityID, &director.CreatedAt, &director.UpdatedAt)
		directors = append(directors, &director)
	}

	return directors, nil
}

func (a *DirectorRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) ([]*domain.Director, error) {
	rows, err := db.Query("SELECT * FROM directors")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed get data from database")
		}
		return nil, err
	}

	var directors []*domain.Director
	for rows.Next() {
		var director domain.Director
		rows.Scan(&director.ID, &director.Name, &director.DateOfBirth, &director.NationalityID, &director.CreatedAt, &director.UpdatedAt)
		directors = append(directors, &director)
	}

	return directors, nil
}

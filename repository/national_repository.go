package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
)

type NationalRepository interface {
	Save(ctx context.Context, tx *sql.Tx, national *domain.National) error
	Update(ctx context.Context, tx *sql.Tx, national *domain.National) error
	Delete(ctx context.Context, tx *sql.Tx, ID int) error
	FindAll(ctx context.Context, db *sql.DB) ([]*domain.National, error)
	FindByName(ctx context.Context, db *sql.DB, name string) (*domain.National, error)
	FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.National, error)
}

type NationalRepositoryaImpl struct {
}

func NewNationalRepository() NationalRepository {
	return &NationalRepositoryaImpl{}
}

func (repository *NationalRepositoryaImpl) Save(ctx context.Context, tx *sql.Tx, national *domain.National) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO national (name) VALUES ($1)", national.Name)
	if err != nil {
		return err
	}

	return nil
}

func (repository *NationalRepositoryaImpl) Update(ctx context.Context, tx *sql.Tx, national *domain.National) error {
	_, err := tx.ExecContext(ctx, "UPDATE national SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", national.Name, national.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *NationalRepositoryaImpl) Delete(ctx context.Context, tx *sql.Tx, ID int) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM national WHERE id = $1", ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New(fmt.Sprintf("Data with ID %d not found", ID))
		}
		return err
	}

	return nil
}

func (repository *NationalRepositoryaImpl) FindAll(ctx context.Context, db *sql.DB) ([]*domain.National, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM national")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed get data from database")
		}
		return nil, err
	}

	var nationals []*domain.National
	for rows.Next() {
		var national domain.National
		rows.Scan(&national.ID, &national.Name, &national.CreatedAt, &national.UpdatedAt)
		nationals = append(nationals, &national)
	}

	return nationals, nil
}

func (repository *NationalRepositoryaImpl) FindByName(ctx context.Context, db *sql.DB, name string) (*domain.National, error) {
	var national domain.National
	err := db.QueryRowContext(ctx, "SELECT * FROM national WHERE name = $1", name).Scan(&national.ID, &national.Name, &national.CreatedAt, &national.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("national with name %s not found", name)
		}
		return nil, err
	}
	return &national, nil
}
func (repository *NationalRepositoryaImpl) FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.National, error) {
	var national domain.National
	err := db.QueryRowContext(ctx, "SELECT * FROM national WHERE id = $1", ID).Scan(&national.ID, &national.Name, &national.CreatedAt, &national.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("national with id %d not found", ID)
		}
		return nil, err
	}

	return &national, nil
}

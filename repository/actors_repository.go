package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
)

type ActorRepository interface {
	Save(ctx context.Context, tx *sql.Tx, actor *domain.Actor) error
	Update(ctx context.Context, tx *sql.Tx, actor *domain.Actor) error
	Delete(ctx context.Context, tx *sql.Tx, ID int) error
	FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Actor, error)
	FindByName(ctx context.Context, db *sql.DB, name string) (*domain.Actor, error)
	FindByNational(ctx context.Context, db *sql.DB, nationalityID int) ([]*domain.Actor, error)
	FindAll(ctx context.Context, db *sql.DB) ([]*domain.Actor, error)
}

type ActorRepositoryImpl struct {
}

func NewActorRepository() ActorRepository {
	return &ActorRepositoryImpl{}
}

func (a *ActorRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, actor *domain.Actor) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO actors (name, date_of_birth, nationality_id) VALUES ($1, $2, $3)", actor.Name, actor.DateOfBirth, actor.NationalityID)
	if err != nil {
		return err
	}

	return nil
}

func (a *ActorRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, actor *domain.Actor) error {
	query := "UPDATE actors SET id = $1, name = $2, date_of_birth = $3, nationality_id = $4, updated_at = $5 WHERE id = $6"
	_, err := tx.ExecContext(ctx, query, actor.ID, actor.Name, actor.DateOfBirth, actor.NationalityID, actor.UpdatedAt, actor.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *ActorRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, ID int) error {
	_, err := tx.Exec("DELETE FROM actors WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *ActorRepositoryImpl) FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Actor, error) {
	var actor domain.Actor
	err := db.QueryRow("SELECT * FROM actors WHERE id = $1", ID).Scan(&actor.ID, &actor.Name, &actor.DateOfBirth, &actor.NationalityID, &actor.CreatedAt, &actor.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("sorry, id not found")
		}
		return nil, err
	}

	return &actor, nil
}

func (a *ActorRepositoryImpl) FindByName(ctx context.Context, db *sql.DB, name string) (*domain.Actor, error) {
	var actor domain.Actor
	err := db.QueryRow("SELECT * FROM actors WHERE name = $1", name).Scan(&actor.ID, &actor.Name, &actor.DateOfBirth, &actor.NationalityID, &actor.CreatedAt, &actor.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("actor with name %s not found", name)
		}
		return nil, err
	}

	return &actor, nil
}

func (a *ActorRepositoryImpl) FindByNational(ctx context.Context, db *sql.DB, nationalityID int) ([]*domain.Actor, error) {
	rows, err := db.Query("SELECT * FROM actors WHERE nationality_id = $1", nationalityID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("nationality with ID %d not found", nationalityID)
		}
		return nil, err
	}

	var actors []*domain.Actor
	for rows.Next() {
		var actor domain.Actor
		rows.Scan(&actor.ID, &actor.Name, &actor.DateOfBirth, &actor.NationalityID, &actor.CreatedAt, &actor.UpdatedAt)
		actors = append(actors, &actor)
	}

	return actors, nil
}

func (a *ActorRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) ([]*domain.Actor, error) {
	rows, err := db.Query("SELECT * FROM actors")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed get data from database")
		}
		return nil, err
	}

	var actors []*domain.Actor
	for rows.Next() {
		var actor domain.Actor
		rows.Scan(&actor.ID, &actor.Name, &actor.DateOfBirth, &actor.NationalityID, &actor.CreatedAt, &actor.UpdatedAt)
		actors = append(actors, &actor)
	}

	return actors, nil
}

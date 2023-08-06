package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var actorReposiory = &repository.ActorRepositoryImpl{}

func testDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestActorService_Save(t *testing.T) {
	db, mock := testDB(t)
	defer db.Close()
	//

	now := time.Now()
	actor := &domain.Actor{
		Name:          "Dimas",
		DateOfBirth:   now,
		NationalityID: 1,
	}

	var actorService = ActorServiceImpl{
		DB:              db,
		ActorRepository: actorReposiory,
	}

	t.Run("Successed Save Data", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO actors").
			WithArgs(actor.Name, actor.DateOfBirth, actor.NationalityID).
			WillReturnResult(sqlmock.NewResult(1, 1)) // You can modify this as needed.
		mock.ExpectCommit()

		// now we execute our method
		err := actorService.Save(context.Background(), &web.ActorModelRequest{
			Name:          "Dimas",
			DateOfBirth:   now,
			NationalityID: 1,
		})
		assert.Nil(t, err)

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Failed Save Data", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO actors").
			WithArgs("", actor.DateOfBirth, 0).
			WillReturnError(errors.New("Data tidak boleh kosong"))
		mock.ExpectRollback()

		// now we execute our method
		err := actorService.Save(context.Background(), &web.ActorModelRequest{
			Name:          "",
			DateOfBirth:   now,
			NationalityID: 0,
		})
		assert.EqualError(t, err, "Data tidak boleh kosong")
	})

}

func recordStats(db *sql.DB, actor *domain.Actor) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(context.Background(), tx)

	actorReposiory.Save(context.Background(), tx, actor)
	return nil
}

package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"log"
)

type RecommendationMovieRepository interface {
	FindAll(ctx context.Context, tx *sql.DB) ([]*domain.RecommendationMovie, error)
	FindByID(ctx context.Context, tx *sql.DB, movieID int) (*domain.RecommendationMovie, error)
	Save(ctx context.Context, tx *sql.Tx, movieID int) error
	Delete(ctx context.Context, tx *sql.Tx, movieID int) error
}

type RecommendationRepositoryImpl struct {
}

func NewRecommendationMovieRepositoryImpl() RecommendationMovieRepository {
	return &RecommendationRepositoryImpl{}
}

func (repository *RecommendationRepositoryImpl) FindAll(ctx context.Context, tx *sql.DB) ([]*domain.RecommendationMovie, error) {
	query := `
			SELECT movie_id AS id,
				   title,
				   release_date,
				   duration,
				   plot,
				   poster_url,
				   trailer_url,
				   language,
				   n.id as national_id,
				   m.created_at as created_at,
				   m.updated_at as updated_at
			FROM recommendation
					 JOIN movies AS m on m.id = recommendation.movie_id
					 JOIN national AS n on n.id = m.nationality_id`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var recommendations []*domain.RecommendationMovie
	for rows.Next() {
		var recommendation domain.RecommendationMovie
		err := rows.
			Scan(&recommendation.ID,
				&recommendation.Title,
				&recommendation.ReleaseDate,
				&recommendation.Duration,
				&recommendation.Plot,
				&recommendation.PosterUrl,
				&recommendation.TrailerUrl,
				&recommendation.Language,
				&recommendation.NationalID,
				&recommendation.CreatedAt,
				&recommendation.UpdatedAt)
		if err != nil {
			return nil, err
		}
		recommendations = append(recommendations, &recommendation)
	}

	return recommendations, nil
}

func (repository *RecommendationRepositoryImpl) FindByID(ctx context.Context, db *sql.DB, movieID int) (*domain.RecommendationMovie, error) {
	query := "SELECT id FROM recommendation WHERE movie_id = $1"
	row := db.QueryRowContext(ctx, query, movieID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var recommendation domain.RecommendationMovie
	err := row.Scan(&recommendation.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &recommendation, nil
}

func (repository *RecommendationRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, movieID int) error {
	query := "INSERT INTO recommendation (movie_id) VALUES ($1)"
	_, err := tx.ExecContext(ctx, query, movieID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *RecommendationRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, movieID int) error {
	query := "DELETE FROM recommendation WHERE movie_id = $1"
	_, err := tx.ExecContext(ctx, query, movieID)
	if err != nil {
		return err
	}

	return nil
}

package services

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
)

type RecommendationMovieService interface {
	FindAll(ctx context.Context) ([]*web.RecommendationMovieModelResponse, error)
	FindByID(ctx context.Context, movieID int) (*web.RecommendationMovieModelResponse, error)
	Save(ctx context.Context, movieID int) error
	Delete(ctx context.Context, movieID int) error
}

type RecommendationMovieServiceImpl struct {
	DB                            *sql.DB
	RecommendationMovieRepository repository.RecommendationMovieRepository
}

func NewRecommendationMovieService(DB *sql.DB, recommendationRepository repository.RecommendationMovieRepository) RecommendationMovieService {
	return &RecommendationMovieServiceImpl{
		DB:                            DB,
		RecommendationMovieRepository: recommendationRepository,
	}
}

func (a *RecommendationMovieServiceImpl) Save(ctx context.Context, movieID int) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	return a.RecommendationMovieRepository.Save(ctx, tx, movieID)
}

func (a *RecommendationMovieServiceImpl) Delete(ctx context.Context, movieID int) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, movieID)
	if err != nil {
		return err
	}

	err = a.RecommendationMovieRepository.Delete(ctx, tx, movieID)
	if err != nil {
		return err
	}

	return nil
}

func (a *RecommendationMovieServiceImpl) FindByID(ctx context.Context, MovieID int) (*web.RecommendationMovieModelResponse, error) {
	result, err := a.RecommendationMovieRepository.FindByID(ctx, a.DB, MovieID)
	if err != nil {
		return nil, err
	}

	return &web.RecommendationMovieModelResponse{
		ID:          result.ID,
		Title:       result.Title,
		ReleaseDate: result.ReleaseDate,
		Duration:    result.Duration,
		Plot:        result.Plot,
		PosterUrl:   result.PosterUrl,
		TrailerUrl:  result.TrailerUrl,
		Language:    result.Language,
		NationalID:  result.NationalID,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (a *RecommendationMovieServiceImpl) FindAll(ctx context.Context) ([]*web.RecommendationMovieModelResponse, error) {
	results, err := a.RecommendationMovieRepository.FindAll(ctx, a.DB)
	if err != nil {
		return nil, err
	}

	var responses []*web.RecommendationMovieModelResponse
	for _, result := range results {
		response := web.RecommendationMovieModelResponse{
			ID:          result.ID,
			Title:       result.Title,
			ReleaseDate: result.ReleaseDate,
			Duration:    result.Duration,
			Plot:        result.Plot,
			PosterUrl:   result.PosterUrl,
			TrailerUrl:  result.TrailerUrl,
			Language:    result.Language,
			NationalID:  result.NationalID,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		}

		responses = append(responses, &response)
	}

	return responses, nil
}

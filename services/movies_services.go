package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
	"time"
)

type MovieService interface {
	Save(ctx context.Context, r *web.MovieModelRequest) (moveiID int, err error)
	Update(ctx context.Context, r *web.MovieModelRequest) error
	Delete(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*web.MovieModelResponse, error)
	FindByTitle(ctx context.Context, name string) (*web.MovieModelResponse, error)
	FindAll(ctx context.Context) ([]*web.MovieModelResponse, error)
	FindAllMoviesByGenreID(ctx context.Context, genreID int) ([]*web.MovieModelResponse, error)
}

type MovieServiceImpl struct {
	DB                   *sql.DB
	MovieRepository      repository.MovieRepository
	movieGenreRepository repository.MovieGenreRepository
}

func NewMovieService(DB *sql.DB, movieRepository repository.MovieRepository) MovieService {
	return &MovieServiceImpl{DB: DB, MovieRepository: movieRepository, movieGenreRepository: repository.NewMovieGenreRepository()}
}

func (service *MovieServiceImpl) Save(ctx context.Context, r *web.MovieModelRequest) (int, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = service.FindByTitle(ctx, r.Title)
	if err == nil {
		return 0, errors.New("movie title already exists")
	}

	releaseDate, err := time.Parse(time.DateOnly, r.ReleaseDate)
	if err != nil {
		return 0, errors.New("incorrect date format yyyy-dd-mm")
	}

	movieID, err := service.MovieRepository.Save(ctx, tx, &domain.Movie{
		Title:       r.Title,
		ReleaseDate: releaseDate,
		Duration:    r.Duration,
		Plot:        r.Plot,
		PosterUrl:   r.PosterUrl,
		TrailerUrl:  r.TrailerUrl,
		Language:    r.Language,
	})
	if err != nil {
		return 0, err
	}

	// Adding data genre_ids after added new movie
	for _, genreID := range r.GenreIDS {
		err := service.movieGenreRepository.Save(ctx, tx, movieID, genreID)
		if err != nil {
			return 0, err
		}
	}

	return movieID, err
}

func (service *MovieServiceImpl) Update(ctx context.Context, r *web.MovieModelRequest) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = service.FindByID(ctx, r.ID)
	if err != nil {
		return err
	}

	// Parsing format date yyyy-mm-dd
	releaseDate, err := time.Parse(time.DateOnly, r.ReleaseDate)
	if err != nil {
		return errors.New("incorrect date format yyyy-mm-dd")
	}

	genresMovie, err := service.movieGenreRepository.FindByID(ctx, service.DB, r.ID)
	if err != nil {
		return err
	}

	for _, genreID := range r.GenreIDS {
		// Check if the genreID exists in the movie's genres
		found := false
		for _, genreIDMovie := range genresMovie.GenreIDS {
			if genreID == genreIDMovie {
				found = true
				break
			}
		}

		// If the genre is not found, will save it
		if !found {
			err := service.movieGenreRepository.Save(ctx, tx, r.ID, genreID)
			if err != nil {
				return err
			}
		}
	}

	// Loop through the movie's genres and check if any need to be deleted
	for _, genreIDMovie := range genresMovie.GenreIDS {
		found := false
		for _, genreID := range r.GenreIDS {
			if genreID == genreIDMovie {
				found = true
				break
			}
		}

		// If the genre is not found in the request, delete it
		if !found {
			err := service.movieGenreRepository.Delete(ctx, tx, r.ID, genreIDMovie)
			if err != nil {
				return err
			}
		}
	}

	return service.MovieRepository.Update(ctx, tx, &domain.Movie{
		ID:          r.ID,
		Title:       r.Title,
		ReleaseDate: releaseDate,
		Duration:    r.Duration,
		Plot:        r.Plot,
		PosterUrl:   r.PosterUrl,
		TrailerUrl:  r.TrailerUrl,
		Language:    r.Language,
	})
}

func (service *MovieServiceImpl) Delete(ctx context.Context, ID int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = service.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	err = service.MovieRepository.Delete(ctx, tx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (service *MovieServiceImpl) FindByID(ctx context.Context, ID int) (*web.MovieModelResponse, error) {
	movieDetail, err := service.MovieRepository.FindByID(ctx, service.DB, ID)
	if err != nil {
		return nil, err
	}

	movieGenre, err := service.movieGenreRepository.FindByID(ctx, service.DB, ID)
	if err != nil {
		return nil, err
	}

	releaseDateFormat := movieDetail.ReleaseDate.Format("2006-01-02")

	return &web.MovieModelResponse{
		ID:          movieDetail.ID,
		Title:       movieDetail.Title,
		ReleaseDate: releaseDateFormat,
		Duration:    movieDetail.Duration,
		Plot:        movieDetail.Plot,
		PosterUrl:   movieDetail.PosterUrl,
		TrailerUrl:  movieDetail.TrailerUrl,
		Language:    movieDetail.Language,
		GenreIDS:    movieGenre.GenreIDS,
		CreatedAt:   movieDetail.CreatedAt,
		UpdatedAt:   movieDetail.UpdatedAt,
	}, nil
}

func (service *MovieServiceImpl) FindByTitle(ctx context.Context, name string) (*web.MovieModelResponse, error) {

	movieDetail, err := service.MovieRepository.FindByTitle(ctx, service.DB, name)
	if err != nil {
		return nil, err
	}

	releaseDateFormat := movieDetail.ReleaseDate.Format("2006-01-02")

	return &web.MovieModelResponse{
		ID:          movieDetail.ID,
		Title:       movieDetail.Title,
		ReleaseDate: releaseDateFormat,
		Duration:    movieDetail.Duration,
		Plot:        movieDetail.Plot,
		PosterUrl:   movieDetail.PosterUrl,
		TrailerUrl:  movieDetail.TrailerUrl,
		Language:    movieDetail.Language,
		CreatedAt:   movieDetail.CreatedAt,
		UpdatedAt:   movieDetail.UpdatedAt,
	}, nil
}

func (service *MovieServiceImpl) FindAll(ctx context.Context) ([]*web.MovieModelResponse, error) {
	moviesDetail, err := service.MovieRepository.FindAll(ctx, service.DB)
	if err != nil {
		return nil, err
	}

	var responses []*web.MovieModelResponse
	for _, movieDetail := range moviesDetail {
		timeFormat := movieDetail.ReleaseDate.Format("2006-01-02")

		genreDetail, _ := service.movieGenreRepository.FindByID(ctx, service.DB, movieDetail.ID)
		response := web.MovieModelResponse{
			ID:          movieDetail.ID,
			Title:       movieDetail.Title,
			ReleaseDate: timeFormat,
			Duration:    movieDetail.Duration,
			Plot:        movieDetail.Plot,
			PosterUrl:   movieDetail.PosterUrl,
			TrailerUrl:  movieDetail.TrailerUrl,
			Language:    movieDetail.Language,
			GenreIDS:    genreDetail.GenreIDS,
			CreatedAt:   movieDetail.CreatedAt,
			UpdatedAt:   movieDetail.UpdatedAt,
		}

		responses = append(responses, &response)
	}

	return responses, nil
}

func (service *MovieServiceImpl) FindAllMoviesByGenreID(ctx context.Context, genreID int) ([]*web.MovieModelResponse, error) {
	moviesDetail, err := service.MovieRepository.FindAllMoviesByGenreID(ctx, service.DB, genreID)
	if err != nil {
		return nil, err
	}

	var responses []*web.MovieModelResponse
	for _, movieDetail := range moviesDetail {

		releaseDateFormat := movieDetail.ReleaseDate.Format("2006-01-02")

		response := web.MovieModelResponse{
			ID:          movieDetail.ID,
			Title:       movieDetail.Title,
			ReleaseDate: releaseDateFormat,
			Duration:    movieDetail.Duration,
			Plot:        movieDetail.Plot,
			PosterUrl:   movieDetail.PosterUrl,
			TrailerUrl:  movieDetail.TrailerUrl,
			Language:    movieDetail.Language,
			CreatedAt:   movieDetail.CreatedAt,
			UpdatedAt:   movieDetail.UpdatedAt,
		}
		responses = append(responses, &response)
	}

	return responses, nil
}

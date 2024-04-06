package sgrpc

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/pb/model"
	"github.com/lancer2672/DandelionServer_Go/pb/request"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetRecentMovies(ctx context.Context, req *request.GetRecentMoviesRequest) (*request.GetRecentMoviesResponse, error) {
	args := db.GetRecentMoviesParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	list, err := server.store.GetRecentMovies(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get recent movies")
	}
	var pbList []*model.Movie
	for _, movie := range list {
		var genres []*model.Genre
		err = json.Unmarshal(movie.Genres, &genres)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse genres")
		}

		var series []*model.Series
		err = json.Unmarshal(movie.Series, &series)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse series")
		}
		pbList = append(pbList, &model.Movie{
			Id:           movie.ID,
			Title:        movie.Title,
			Duration:     movie.Duration,
			Description:  movie.Description,
			ActorAvatars: movie.ActorAvatars,
			Trailer:      movie.Trailer,
			FilePath:     movie.FilePath,
			Thumbnail:    movie.Thumbnail,
			Views:        movie.Views,
			Genres:       genres,
			Series:       series,
			Stars:        movie.Stars,
			CreatedAt:    timestamppb.New(movie.CreatedAt),
		})
	}
	response := &request.GetRecentMoviesResponse{Data: pbList}
	return response, nil
}
func (server *Server) SearchMovies(ctx context.Context, req *request.SearchMoviesRequest) (*request.SearchMoviesResponse, error) {
	args := db.SearchMoviesParams{
		Column1: sql.NullString{String: req.GetColumn_1(), Valid: true},
		Limit:   req.GetLimit(),
		Offset:  req.GetOffset(),
	}
	list, err := server.store.SearchMovies(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to search movies")
	}
	var pbList []*model.Movie
	for _, movie := range list {
		// Parse Genres and Series from JSON
		var genres []*model.Genre
		err = json.Unmarshal(movie.Genres, &genres)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse genres")
		}

		var series []*model.Series
		err = json.Unmarshal(movie.Series, &series)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse series")
		}

		pbList = append(pbList, &model.Movie{
			Id:           movie.ID,
			Title:        movie.Title,
			Duration:     movie.Duration,
			Description:  movie.Description,
			ActorAvatars: movie.ActorAvatars,
			Trailer:      movie.Trailer,
			FilePath:     movie.FilePath,
			Thumbnail:    movie.Thumbnail,
			Views:        movie.Views,
			Stars:        movie.Stars,
			CreatedAt:    timestamppb.New(movie.CreatedAt),
			Genres:       genres,
			Series:       series,
		})
	}
	response := &request.SearchMoviesResponse{Data: pbList}
	return response, nil
}

func (server *Server) CreateMovie(ctx context.Context, req *request.CreateMovieRequest) (*model.Movie, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {

		return nil, status.Errorf(codes.Internal, "Failed to get metadata")
	}
	fmt.Println("Metadata", md)
	args := db.CreateMovieParams{
		Title:        req.GetTitle(),
		Duration:     req.GetDuration(),
		Description:  req.GetDescription(),
		ActorAvatars: req.GetActorAvatars(),
		Trailer:      req.GetTrailer(),
		FilePath:     req.GetFilePath(),
		Thumbnail:    req.GetThumbnail(),
		Views:        req.GetViews(),
		Stars:        req.GetStars(),
	}
	_, err := server.store.CreateMovie(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create movie")
	}
	movie := &model.Movie{
		Title:        args.Title,
		Duration:     args.Duration,
		Description:  args.Description,
		ActorAvatars: args.ActorAvatars,
		Trailer:      args.Trailer,
		FilePath:     args.FilePath,
		Thumbnail:    args.Thumbnail,
		Views:        args.Views,
		Stars:        args.Stars,
		CreatedAt:    timestamppb.Now(), // Assuming the movie is created now
	}

	return movie, nil
}

func (server *Server) GetListMovies(ctx context.Context, req *request.GetListMoviesRequest) (*request.GetListMoviesResponse, error) {
	args := db.GetListMoviesParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	list, err := server.store.GetListMovies(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get list of movies")
	}
	var pbList []*model.Movie
	for _, movie := range list {
		var genres []*model.Genre
		err = json.Unmarshal(movie.Genres, &genres)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse genres")
		}

		var series []*model.Series
		err = json.Unmarshal(movie.Series, &series)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse series")
		}

		pbList = append(pbList, &model.Movie{
			Id:           movie.ID,
			Title:        movie.Title,
			Duration:     movie.Duration,
			Description:  movie.Description,
			ActorAvatars: movie.ActorAvatars,
			Trailer:      movie.Trailer,
			FilePath:     movie.FilePath,
			Thumbnail:    movie.Thumbnail,
			Views:        movie.Views,
			Stars:        movie.Stars,
			Genres:       genres,
			Series:       series,
			CreatedAt:    timestamppb.New(movie.CreatedAt),
		})
	}
	response := &request.GetListMoviesResponse{Data: pbList}
	return response, nil
}

func (server *Server) GetMovie(ctx context.Context, req *request.GetMovieRequest) (*model.Movie, error) {
	movie, err := server.store.GetMovie(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get movie")
	}
	var genres []*model.Genre
	err = json.Unmarshal(movie.Genres, &genres)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to parse genres")
	}

	var series []*model.Series
	err = json.Unmarshal(movie.Series, &series)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to parse series")
	}
	pbMovie := &model.Movie{
		Id:           movie.ID,
		Title:        movie.Title,
		Duration:     movie.Duration,
		Description:  movie.Description,
		ActorAvatars: movie.ActorAvatars,
		Trailer:      movie.Trailer,
		FilePath:     movie.FilePath,
		Thumbnail:    movie.Thumbnail,
		Views:        movie.Views,
		Stars:        movie.Stars,
		Genres:       genres,
		Series:       series,
		CreatedAt:    timestamppb.New(movie.CreatedAt),
	}
	return pbMovie, nil
}
func (server *Server) UpdateMovie(ctx context.Context, req *request.UpdateMovieRequest) (*request.UpdateMovieResponse, error) {
	args := db.UpdateMovieParams{
		Title: sql.NullString{
			String: req.GetTitle(),
			Valid:  req.GetTitle() != "",
		},
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.GetDescription() != "",
		},
		FilePath: sql.NullString{
			String: req.GetFilePath(),
			Valid:  req.GetFilePath() != "",
		},
		Thumbnail: sql.NullString{
			String: req.GetThumbnail(),
			Valid:  req.GetThumbnail() != "",
		},
		ID: req.GetId(),
	}
	fmt.Println(args)
	movie, err := server.store.UpdateMovie(ctx, args)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update movie")
		return nil, status.Errorf(codes.Internal, "Failed to update movie")
	}
	pbMovie := &model.Movie{
		Id:           movie.ID,
		Title:        movie.Title,
		Duration:     movie.Duration,
		Description:  movie.Description,
		ActorAvatars: movie.ActorAvatars,
		Trailer:      movie.Trailer,
		FilePath:     movie.FilePath,
		Thumbnail:    movie.Thumbnail,
		Views:        movie.Views,
		Stars:        movie.Stars,

		CreatedAt: timestamppb.New(movie.CreatedAt),
	}
	return &request.UpdateMovieResponse{
		Movie: pbMovie,
	}, nil
}

func (server *Server) GetMoviesByGenre(ctx context.Context, req *request.GetMoviesByGenreRequest) (*request.GetMoviesByGenreResponse, error) {
	args := db.GetMoviesByGenreParams{
		GenreID: req.GetGenreId(),
		Limit:   req.GetLimit(),
		Offset:  req.GetOffset(),
	}
	list, err := server.store.GetMoviesByGenre(ctx, args)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get movies by genre")
		return nil, status.Errorf(codes.Internal, "Failed to get movies by genre")
	}

	var pbList []*model.Movie
	for _, movie := range list {
		var genres []*model.Genre
		err = json.Unmarshal(movie.Genres, &genres)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse genres")
		}

		var series []*model.Series
		err = json.Unmarshal(movie.Series, &series)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse series")
		}

		pbList = append(pbList, &model.Movie{
			Id:           movie.ID,
			Title:        movie.Title,
			Duration:     movie.Duration,
			Description:  movie.Description,
			ActorAvatars: movie.ActorAvatars,
			Trailer:      movie.Trailer,
			FilePath:     movie.FilePath,
			Thumbnail:    movie.Thumbnail,
			Genres:       genres,
			Series:       series,
			Views:        movie.Views,
			Stars:        movie.Stars,
			CreatedAt:    timestamppb.New(movie.CreatedAt),
		})
	}
	response := &request.GetMoviesByGenreResponse{Data: pbList}
	return response, nil
}

func (server *Server) GetMoviesBySerie(ctx context.Context, req *request.GetMoviesBySerieRequest) (*request.GetMoviesBySerieResponse, error) {
	args := db.GetMoviesBySerieParams{
		SeriesID: req.GetId(),
		Limit:    req.GetLimit(),
		Offset:   req.GetOffset(),
	}
	list, err := server.store.GetMoviesBySerie(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get movies by series")
	}
	var pbList []*model.Movie
	for _, movie := range list {
		var genres []*model.Genre
		err = json.Unmarshal(movie.Genres, &genres)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse genres")
		}

		var series []*model.Series
		err = json.Unmarshal(movie.Series, &series)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse series")
		}

		pbList = append(pbList, &model.Movie{
			Id:           movie.ID,
			Title:        movie.Title,
			Duration:     movie.Duration,
			Description:  movie.Description,
			ActorAvatars: movie.ActorAvatars,
			Trailer:      movie.Trailer,
			FilePath:     movie.FilePath,
			Thumbnail:    movie.Thumbnail,
			Genres:       genres,
			Series:       series,
			Views:        movie.Views,
			Stars:        movie.Stars,
			CreatedAt:    timestamppb.New(movie.CreatedAt),
		})
	}
	response := &request.GetMoviesBySerieResponse{Data: pbList}
	return response, nil
}
func (server *Server) GetWatchingMovies(ctx context.Context, req *request.GetWatchingMoviesRequest) (*request.GetWatchingMoviesResponse, error) {
	args := db.GetWatchingMoviesParams{
		UserID: req.GetUserId(),
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	list, err := server.store.GetWatchingMovies(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get watching movies")
	}
	var pbList []*model.WatchingMovie
	for _, movie := range list {

		var genres []*model.Genre
		err = json.Unmarshal(movie.Genres, &genres)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse genres")
		}

		var series []*model.Series
		err = json.Unmarshal(movie.Series, &series)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse series")
		}
		pbList = append(pbList, &model.WatchingMovie{
			Movie: &model.Movie{
				Id:           movie.ID,
				Title:        movie.Title,
				Duration:     movie.Duration,
				Description:  movie.Description,
				ActorAvatars: movie.ActorAvatars,
				Trailer:      movie.Trailer,
				FilePath:     movie.FilePath,
				Thumbnail:    movie.Thumbnail,
				Views:        movie.Views,
				Stars:        movie.Stars,
				Genres:       genres,
				Series:       series,
				CreatedAt:    timestamppb.New(movie.CreatedAt),
			},
			WatchedDuration: float32(movie.WatchedDuration),
			LastWatched:     timestamppb.New(movie.LastWatched),
		})
	}
	response := &request.GetWatchingMoviesResponse{Data: pbList}
	return response, nil
}

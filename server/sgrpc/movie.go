package sgrpc

import (
	"context"
	"database/sql"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/pb/model"
	"github.com/lancer2672/DandelionServer_Go/pb/request"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetRecentMovies(ctx context.Context, req *request.GetRecentMoviesRequest) (*request.GetRecentMoviesResponse, error) {
	list, err := server.store.GetRecentMovies(ctx, req.GetLimit())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get recent movies")
	}
	var pbList []*model.Movie
	for _, movie := range list {
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
		})
	}
	response := &request.GetRecentMoviesResponse{Data: pbList}
	return response, nil
}

func (server *Server) SearchMovies(ctx context.Context, req *request.SearchMoviesRequest) (*request.SearchMoviesResponse, error) {
	args := db.SearchMoviesParams{
		Column1: sql.NullString{String: req.GetColumn1(), Valid: true},
		Limit:   req.GetLimit(),
		Offset:  req.GetOffset(),
	}
	list, err := server.store.SearchMovies(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to search movies")
	}
	var pbList []*model.Movie
	for _, movie := range list {
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
		})
	}
	response := &request.SearchMoviesResponse{Data: pbList}
	return response, nil
}
func (server *Server) CreateMovie(ctx context.Context, req *request.CreateMovieRequest) (*model.Movie, error) {
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
	err := server.store.CreateMovie(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create movie")
	}

	// Assuming the store returns the created movie
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
		CreatedAt:    timestamppb.New(movie.CreatedAt),
	}
	return pbMovie, nil
}

func (server *Server) GetMoviesByGenre(ctx context.Context, req *request.GetMoviesByGenreRequest) (*request.GetMoviesByGenreResponse, error) {
	args := db.GetMoviesByGenreParams{
		GenreID: req.GetGenreId(),
		Limit:   req.GetLimit(),
		Offset:  req.GetOffset(),
	}
	list, err := server.store.GetMoviesByGenre(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get movies by genre")
	}
	var pbList []*model.Movie
	for _, movie := range list {
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
		})
	}
	response := &request.GetMoviesByGenreResponse{Data: pbList}
	return response, nil
}

func (server *Server) GetMoviesBySerie(ctx context.Context, req *request.GetMoviesBySerieRequest) (*request.GetMoviesBySerieResponse, error) {
	args := db.GetMoviesBySerieParams{
		ID:     req.GetId(),
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	list, err := server.store.GetMoviesBySerie(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get movies by series")
	}
	var pbList []*model.Movie
	for _, movie := range list {
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
		})
	}
	response := &request.GetMoviesBySerieResponse{Data: pbList}
	return response, nil
}

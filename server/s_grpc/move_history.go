package sgrpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/pb/model"
	"github.com/lancer2672/DandelionServer_Go/pb/request"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateMovieHistory(ctx context.Context, req *request.CreateMovieHistoryRequest) (*empty.Empty, error) {

	args := db.CreateMovieHistoryParams{
		UserID:          req.GetUserId(),
		MovieID:         req.GetMovieId(),
		WatchedDuration: req.GetWatchedDuration(),
		LastWatched:     req.GetLastWatched().AsTime(),
	}
	err := server.store.CreateMovieHistory(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Create movie histor failed")
	}

	return &empty.Empty{}, nil
}
func (server *Server) GetMovieHistory(ctx context.Context, req *request.GetMovieHistoryRequest) (*request.GetMovieHistoryResponse, error) {

	list, err := server.store.GetMovieHistoryByUserId(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Create movie histor failed")
	}
	var pbList []*model.MovieHistory
	for _, movie := range list {
		pbList = append(pbList, &model.MovieHistory{
			MovieId:         movie.MovieID,
			UserId:          movie.UserID,
			WatchedDuration: movie.WatchedDuration,
			LastWatched:     timestamppb.New(movie.LastWatched),
		})
	}

	// Create the response message with the converted list
	response := &request.GetMovieHistoryResponse{Data: pbList}
	return response, nil
}

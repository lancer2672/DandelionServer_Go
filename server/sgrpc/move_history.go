package sgrpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateMovieHistory(ctx context.Context, req *pb.CreateMovieHistoryRequest) (*empty.Empty, error) {

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

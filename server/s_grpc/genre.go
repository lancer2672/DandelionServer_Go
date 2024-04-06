package sgrpc

import (
	"context"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
	"github.com/lancer2672/DandelionServer_Go/pb/model"
	"github.com/lancer2672/DandelionServer_Go/pb/request"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *MovieService) GetListGenres(ctx context.Context, req *request.GetListGenresRequest) (*request.GetListGenresResponse, error) {
	args := db.GetListGenresParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	list, err := server.store.GetListGenres(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get list of genres")
	}
	var pbList []*model.Genre
	for _, genre := range list {
		pbList = append(pbList, &model.Genre{
			Id:   genre.ID,
			Name: genre.Name,
		})
	}
	response := &request.GetListGenresResponse{Data: pbList}
	return response, nil
}

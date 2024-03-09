package service

import (
	"context"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
)

type UploadService struct {
	store *db.Store
}

type UploadServiceInterface interface {
	UploadFile(ctx context.Context) error
}

func NewUploadService(store *db.Store) *UploadService {
	return &UploadService{
		store: store,
	}
}

func (s *UploadService) UploadFile(ctx context.Context) (error) {
	// fmt.Println(arg)
	// if err != nil {
	// 	return err
	// }
	return  nil
}

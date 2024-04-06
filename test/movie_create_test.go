package test

import (
	"context"
	"testing"

	"github.com/lancer2672/DandelionServer_Go/pb/request"
	"github.com/lancer2672/DandelionServer_Go/pb/service"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestCreateMovie(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	client := service.NewMovieServiceClient(conn)
	req := &request.CreateMovieRequest{
		Title:        "Test Movie",
		Duration:     120,
		Description:  "This is a test movie",
		ActorAvatars: []string{"actor1.jpg", "actor2.jpg"},
		Trailer:      "trailer.mp4",
		FilePath:     "movie.mp4",
		Thumbnail:    "thumbnail.jpg",
		Views:        0,
		Stars:        5,
	}

	movie, err := client.CreateMovie(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, movie)

	require.Equal(t, req.GetTitle(), movie.Title)
	require.Equal(t, req.GetDuration(), movie.Duration)
	require.Equal(t, req.GetDescription(), movie.Description)
	require.Equal(t, req.GetActorAvatars(), movie.ActorAvatars)
	require.Equal(t, req.GetTrailer(), movie.Trailer)
	require.Equal(t, req.GetFilePath(), movie.FilePath)
	require.Equal(t, req.GetThumbnail(), movie.Thumbnail)
	require.Equal(t, req.GetViews(), movie.Views)
	require.Equal(t, req.GetStars(), movie.Stars)
}

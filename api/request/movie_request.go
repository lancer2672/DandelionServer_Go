package requests

import "database/sql"

type MoviesByGenreRequest struct {
	GenreID int32 `uri:"id" binding:"required"`
	Limit   int64 `form:"limit,default=10"`
	Offset  int64 `form:"offset"`
}

type MoviesBySeriesRequest struct {
	ID     int32 `uri:"id"  binding:"required"`
	Limit  int64 `form:"limit,default=10"`
	Offset int64 `form:"offset"`
}

type SearchMoviesRequest struct {
	Column1 sql.NullString `form:"column_1"`
	Limit   int64          `form:"limit,default=10"`
	Offset  int64          `form:"offset"`
}

type CreateMovieRequest struct {
	Title        string   `json:"title" binding:"required"`
	Duration     int32    `json:"duration" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	ActorAvatars []string `json:"actor_avatars" binding:"required"`
	Trailer      string   `json:"trailer" binding:"required"`
	FilePath     string   `json:"file_path" binding:"required"`
	Thumbnail    string   `json:"thumbnail" binding:"required"`
	Views        int32    `json:"views" binding:"required"`
	Stars        int32    `json:"stars" binding:"required"`
}
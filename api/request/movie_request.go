package requests

type MoviesByGenreRequest struct {
	GenreID int32 `form:"genre_id" binding:"required"`
	Limit   int64 `form:"limit,default=10"`
	Offset  int64 `form:"offset"`
}

type MoviesBySeriesRequest struct {
	ID     int32 `form:"series_id"  binding:"required"`
	Limit  int64 `form:"limit,default=10"`
	Offset int64 `form:"offset"`
}

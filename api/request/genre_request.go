package requests

type GetListGenresRequest struct {
	Limit  int64 `form:"limit,default=10"`
	Offset int64 `form:"offset"`
}

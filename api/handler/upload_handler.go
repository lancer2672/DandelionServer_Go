package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
}

type UploadHandlerInterface interface {
	UploadFile(ctx *gin.Context)
}

func NewUploadHandler() UploadHandlerInterface {
	return &UploadHandler{}
}
func (handler *UploadHandler) UploadFile(ctx *gin.Context) {
	// single file
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	ctx.SaveUploadedFile(file, "./upload")

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

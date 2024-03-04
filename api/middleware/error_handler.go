package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lancer2672/DandelionServer_Go/utils"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("ERROR", err)
				c.JSON(http.StatusInternalServerError, utils.FormatErrorRes(err.(error)))
			}
		}()

		c.Next()
	}
}

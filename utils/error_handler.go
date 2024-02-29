package utils

import "github.com/gin-gonic/gin"

func ErrorFormatter(err error) gin.H {
	return gin.H{"error": err.Error()}
}

package handlers

import "github.com/gin-gonic/gin"

func newResponseErrBindJSON(err error) *gin.H {
	return &gin.H{
		"error": err.Error(),
	}
}

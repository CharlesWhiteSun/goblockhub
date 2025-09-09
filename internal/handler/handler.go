package handler

import "github.com/gin-gonic/gin"

type IPlatformHandler interface {
	RegisterRoutes(r *gin.Engine)
}

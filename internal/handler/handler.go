package handler

import "github.com/gin-gonic/gin"

// 所有平台 Handler 都實作這個介面
type IPlatformHandler interface {
	RegisterRoutes(r *gin.Engine)
}

package router

import (
	"github.com/choisangh/gin_project/pkg/api"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/images", api.CreateImage)
	r.GET("/images/:id", api.ReadImage)

	return r
}

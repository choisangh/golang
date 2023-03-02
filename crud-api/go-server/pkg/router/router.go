package router

import (
	"github.com/choisangh/crud-api/go-server/pkg/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	crudRouter := r.Group("/api/cruds")

	crudRouter.GET("", apis.GetBoardList)
	crudRouter.POST("", apis.CreateBoard)
	crudRouter.GET("/:id", apis.GetBoardByID)
	crudRouter.PATCH("/:id", apis.UpdateBoard)
	crudRouter.DELETE("/:id", apis.DeleteBoardByID)

	return r
}

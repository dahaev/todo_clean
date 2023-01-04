package handler

import (
	"github.com/dahaev/todo.git/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Services
}

func NewHandler(service *services.Services) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIndentity)
	{
		lists := api.Group("lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.GotAllList)
			lists.GET("/:id", h.GetListByID)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.GotAllItem)
				items.GET("/:items_id", h.GetItemByID)
				items.PUT("/:items_id", h.UpdateItem)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}
	}
	return router
}

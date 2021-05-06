package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zemags/go_workshop_2/pkg/service"
)

type Handler struct {
	// dependency injection by clean archirecture
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Register and authorize endpoints
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// Work with todo lists
	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.allLists)
			// : - means any value, can access by id
			lists.GET("/:id", h.listByID)
			lists.PUT("/:id", h.updateList)
			lists.DELETE(":id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.allItems)
				items.GET("/:item_id", h.itemByID)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE(":item_id", h.deleteItem)
			}
		}
	}

	return router
}

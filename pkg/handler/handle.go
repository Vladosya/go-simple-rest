package handler

import (
	"github.com/Vladosya/go-test-rest/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler { // создаём новый handler с полем services
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine { // обработчик роутов, Создание роутов
	router := gin.New() // инициализация роутов

	api := router.Group("/api-v1")
	{
		api.POST("/user", h.createUser)
		api.GET("/user", h.getUsers)
		api.GET("/user/:id", h.getUserById)
		api.PUT("/user/:id", h.updateUser)
		api.DELETE("/user/:id", h.deleteUserById)

		api.POST("/post", h.createPost)
		api.GET("/post", h.getPosts)
		api.DELETE("/post/:id", h.deletePostById)
	}

	return router
}

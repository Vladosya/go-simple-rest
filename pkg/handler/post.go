package handler

import (
	todo "github.com/Vladosya/go-test-rest"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createPost(c *gin.Context) {
	var body todo.Post
	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreatePost(body)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"id":      id,
		"message": "Успешное создание поста",
	})
}

func (h *Handler) getPosts(c *gin.Context) {
	posts, err := h.services.GetPosts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(posts) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Успешное получение постов",
			"result":  posts,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Список постов пуст",
			"result":  []int{},
		})
	}
}

func (h *Handler) deletePostById(c *gin.Context) {
	needId, res := c.Params.Get("id")
	if res != true || !govalidator.IsInt(needId) {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Плохой запрос. Измените данные и попробуйте ещё раз",
		})
		return
	}
	postDeleted, err := h.services.DeletePostById(needId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(postDeleted) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Успешное удаление поста",
			"result":  postDeleted,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Пост с таким id не существует",
			"result":  []int{},
		})
	}
}

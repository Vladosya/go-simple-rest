package handler

import (
	"fmt"
	todo "github.com/Vladosya/go-test-rest"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	var body todo.User
	if err := c.BindJSON(&body); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(body)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"id":      id,
		"message": "Успешное создание пользователя",
	})
}

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.services.GetUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(users) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Успешное получение пользователей",
			"result":  users,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Список пользователей пуст",
			"result":  []int{},
		})
	}
}

func (h *Handler) getUserById(c *gin.Context) {
	needId, res := c.Params.Get("id")
	if res != true || !govalidator.IsInt(needId) {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Плохой запрос. Измените данные и попробуйте ещё раз",
		})
		return
	}
	user, err := h.services.GetUserById(needId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(user) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Успешное получение пользователя",
			"result":  user,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Пользователь с таким id не существует",
			"result":  []int{},
		})
	}
}

func (h *Handler) updateUser(c *gin.Context) {
	fmt.Println("updateUser")
}

func (h *Handler) deleteUserById(c *gin.Context) {
	needId, res := c.Params.Get("id")
	if res != true || !govalidator.IsInt(needId) {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Плохой запрос. Измените данные и попробуйте ещё раз",
		})
		return
	}
	userDeleted, err := h.services.DeleteUserById(needId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(userDeleted) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Успешное удаление пользователя",
			"result":  userDeleted,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Пользователь с таким id не существует",
			"result":  []int{},
		})
	}
}

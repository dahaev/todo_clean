package handler

import (
	"net/http"

	"github.com/dahaev/todo.git"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	var input todo.Todolist
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) GotAllList(c *gin.Context) {

}

func (h *Handler) GetListByID(c *gin.Context) {

}

func (h *Handler) UpdateList(c *gin.Context) {

}

func (h *Handler) DeleteList(c *gin.Context) {

}

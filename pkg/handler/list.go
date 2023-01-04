package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userCTX)
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

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

// signUp - register user in system
func (h *Handler) signUp(c *gin.Context) {
	// input - user struct
	var input workshop_2.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}

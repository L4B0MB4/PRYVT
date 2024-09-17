package controller

import (
	"net/http"
	"strings"

	"github.com/L4B0MB4/PRYVT/identification/pkg/command/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewEventController() *UserController {
	return &UserController{}
}

func (ctrl *UserController) ChangeName(c *gin.Context) {

	userId := c.Param("userId")

	if len(strings.TrimSpace(userId)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Path param cant be empty or null"})
		return
	}
	var m models.ChangeName
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

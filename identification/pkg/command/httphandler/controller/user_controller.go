package controller

import (
	"net/http"
	"strings"

	"github.com/L4B0MB4/PRYVT/identification/pkg/command/aggregates"
	"github.com/L4B0MB4/PRYVT/identification/pkg/command/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (ctrl *UserController) ChangeName(c *gin.Context) {

	userId := c.Param("userId")

	if len(strings.TrimSpace(userId)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Path param cant be empty or null"})
		return
	}
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var m models.ChangeName
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ua, err := aggregates.NewUserAggregate(userUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = ua.ChangeName(m.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

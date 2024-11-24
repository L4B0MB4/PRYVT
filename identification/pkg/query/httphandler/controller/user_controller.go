package controller

import (
	"net/http"
	"strings"

	"github.com/L4B0MB4/PRYVT/identification/pkg/aggregates"
	"github.com/L4B0MB4/PRYVT/identification/pkg/query/models"
	"github.com/L4B0MB4/PRYVT/identification/pkg/query/store/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userRepo *repository.UserRepository
}

func NewUserController(userRepo *repository.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

func (ctrl *UserController) GetUser(c *gin.Context) {

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

	ua, err := aggregates.NewUserAggregate(userUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(ua.Events) > 0 {
		m := models.UserInfo{
			DisplayName: ua.DisplayName,
			Name:        ua.Name,
			Email:       ua.Email,
			ChangeDate:  ua.ChangeDate,
		}
		c.JSON(http.StatusOK, m)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

}

func (ctrl *UserController) GetUsers(c *gin.Context) {

}

func (ctrl *UserController) GetUserFromDB(c *gin.Context) {
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

	user, err := ctrl.userRepo.GetUserById(userUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

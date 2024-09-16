package controller

import (
	"net/http"
	"strings"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/models"
	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/models/customerrors"
	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/store"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	repo *store.EventRepository
}

func NewEventController(repo *store.EventRepository) *EventController {
	return &EventController{
		repo: repo,
	}
}

func (ctrl *EventController) GetEventsForAggregate(c *gin.Context) {

	aggregateId := c.Param("aggregateId")

	if len(strings.TrimSpace(aggregateId)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Path param cant be empty or null"})
		return
	}

	resp, err := ctrl.repo.GetEventsForAggregate(aggregateId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unkown error occured"})
		return
	}
	if len(resp) == 0 {
		c.JSON(http.StatusOK, []models.Event{})
		return
	}
	c.JSON(http.StatusOK, &resp)
}

func (ctrl *EventController) AddEventToAggregate(c *gin.Context) {
	var event models.Event
	aggregateId := c.Param("aggregateId")
	if len(strings.TrimSpace(aggregateId)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Path param cant be empty or null"})
		return
	}
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.AggregateId = aggregateId
	err := ctrl.repo.AddEvent(&event)
	if err != nil {
		_, ok := err.(*customerrors.DuplicateVersionError)
		if ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error trying to add the same event multiple times"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unkown error occured"})
		return
	}
}

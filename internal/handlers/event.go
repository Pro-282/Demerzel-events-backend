package handlers

import (
	"demerzel-events/internal/db"
	"demerzel-events/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEventHandler(c *gin.Context) {
	var input models.NewEvent

	// Error if JSON request is invalid
	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO use auth middleware to get creator of event.
	if input.CreatorId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User id could not be found"})
		return
	}

	createdEvent, err := models.CreateEvent(db.DB, &input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Event Created": createdEvent})

}
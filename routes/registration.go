package routes

import (
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(c *gin.Context) {
	userId := c.GetInt64("user_id")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to register for event"})
		return
	}

	c.JSON(200, gin.H{"message": "Successfully registered for event"})
}
func UnregisterFromEvent(c *gin.Context) {
	userId := c.GetInt64("user_id")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}
	err = event.Unregister(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to unregister from event"})
		return
	}

	c.JSON(200, gin.H{"message": "Successfully unregistered from event"})
}

package routes

import (
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve events"})
		return
	}
	c.JSON(200, events)
}

func GetEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(200, event)
}

func CreateEvent(c *gin.Context) {
	var newEvent models.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetInt64("user_id")
	newEvent.UserID = userId
	if err := newEvent.Save(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, newEvent)
}

func UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}
	var event *models.Event
	event, err = models.GetEventByID(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}
	userId := c.GetInt64("user_id")
	if event.UserID != userId {
		c.JSON(403, gin.H{"error": "You are not authorized to update this event"})
		return
	}
	var updatedEvent models.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updatedEvent.ID = id
	if err := updatedEvent.UpdateEvent(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update event"})
		return
	}
	c.JSON(200, updatedEvent)
}

func DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(id)

	userId := c.GetInt64("user_id")
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userId {
		c.JSON(403, gin.H{"error": "You are not authorized to delete this event"})
		return
	}

	if err := event.DeleteEvent(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete event"})
		return
	}
	c.JSON(204, nil)
}

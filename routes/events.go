package routes

import (
	"net/http"
	"strconv"

	"example.com/restAPIProject/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not fetch events. Try again later.", "error": err})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messege": "Could not parse event id.", "error": err})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not fetch event. Try again later.", "error": err})
		return
	}

	context.JSON(http.StatusOK, event)
}

func creatEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messege": "Could not parse request data.", "error": err})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not create events. Try again later.", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"messege": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messege": "Could not parse event id.", "error": err})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not fetch the event. Try again later.", "error": err})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messege": "Could not parse request data.", "error": err})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not update the event. Try again later.", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"messege": "Event updated successfully."})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messege": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not fetch the event. Try again later."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not delete the event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"messege": "Event deleted successfully."})
}

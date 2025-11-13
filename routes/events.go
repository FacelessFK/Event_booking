package routes

import (
	"net/http"
	"strconv"

	"github.com/FacelessFK/Event_booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}
	ctx.JSON(http.StatusOK, event)

}

func createEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	}

	event.UserID = userId

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later."})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	var updateEvent models.Event
	err = ctx.ShouldBindJSON(&updateEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again later."})
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event updated successfully!", "id": eventId})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event. Try again later."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted successfully!"})
}

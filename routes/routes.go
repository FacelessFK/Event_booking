package routes

import (
	"github.com/FacelessFK/Event_booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	// server.POST("/events", middlewares.Authenticate, createEvent)

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", createEvent)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}

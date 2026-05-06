package routes

import (
	"example.com/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Events Endpoints
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)

	// Authenticated is a group of endpoints that needs an user token
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("events/:eventId", updateEvent)
	authenticated.DELETE("events/:eventId", deleteEvent)

	// Registrations Endpoints
	authenticated.POST("/events/:eventId/register", registerForEvent)
	authenticated.DELETE("/events/:eventId/register", cancelRegistration)

	// Users Endpoints
	server.POST("/signup", signup)
	server.POST("/login", login)

}

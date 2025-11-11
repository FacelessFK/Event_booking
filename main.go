package main

import (
	"github.com/FacelessFK/Event_booking/db"
	"github.com/FacelessFK/Event_booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

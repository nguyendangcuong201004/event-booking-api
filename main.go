package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyendangcuong201004/event-booking-api/db"
	"github.com/nguyendangcuong201004/event-booking-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

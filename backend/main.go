package main

import (
    "ai-task-manager/auth"
    "ai-task-manager/database"
    "ai-task-manager/websocket"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Initialize Database
    database.InitDB()

    // Routes
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.POST("/login", auth.Login)
    r.GET("/tasks", database.GetTasks)
    r.POST("/tasks", database.CreateTask)

    // WebSocket Endpoint
    r.GET("/ws", websocket.HandleConnections)

    r.Run(":8080")
}

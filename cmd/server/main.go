package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}

/*

gin-rest-api/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   └── routes/
│
├── middleware/
│
├── go.mod
└── go.sum


server
   ↓
HTTP API

worker
   ↓
Background Jobs

migration
   ↓
Database Migration


*/

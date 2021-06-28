package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()

	router := gin.New()

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run("localhost:8080")
}

package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Message struct {
	SQL string `json:sql`
}

func main() {
	addr := ":8080"
	r := gin.New()
	r.POST("/metrics", func(c *gin.Context) {
		msg := &Message{}
		// Send raw sql message
		if err := c.ShouldBind(msg); err == nil {
			log.Println(msg.SQL)
		}
	})
	r.Run(addr)
}

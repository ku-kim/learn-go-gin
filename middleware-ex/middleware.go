package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Custom mildderware : https://gin-gonic.com/ko-kr/docs/examples/custom-middleware/
func CustomLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", 12345)
		time.Sleep(3 * time.Second)

		c.Next() // Request 기준
		// Request 이후
		latency := time.Since(t)
		log.Print(latency)

		// 송싱할 상태 코드에 접근
		status := c.Writer.Status()
		log.Println(status)
	}
}

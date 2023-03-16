package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Custom mildderware : https://gin-gonic.com/ko-kr/docs/examples/custom-middleware/
func CustomLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Request 이전
		t := time.Now()
		c.Set("example", 12345)
		log.Print("c.Next 이전, 2초 대기 . . .")
		time.Sleep(2 * time.Second)
		log.Print("2초 끝")

		c.Next() // Request 기준

		// Request 이후
		log.Print("c.next 이후, 3초 대기 . .")
		latency := time.Since(t)
		time.Sleep(3 * time.Second)
		log.Print(latency)
		log.Print("3초 끝")

		// response status code 접근
		status := c.Writer.Status()
		log.Println(status)
	}
}

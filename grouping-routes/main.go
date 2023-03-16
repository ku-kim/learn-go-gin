package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// https://gin-gonic.com/ko-kr/docs/examples/grouping-routes/
func main() {
	router := gin.Default()

	// v1 그룹
	v1 := router.Group("/v1")
	{
		v1.GET("/login", hello)
		v1.GET("/submit", hello)
		v1.GET("/read", hello)
	}

	// v2 그룹
	v2 := router.Group("/v2")
	{
		v2.GET("/login", hello)
		v2.GET("/submit", hello)
		v2.GET("/read", hello)
	}

	router.Run(":8080")
}

func hello(c *gin.Context) {
	log.Println("Hello")
	c.IndentedJSON(http.StatusOK, "world")
}

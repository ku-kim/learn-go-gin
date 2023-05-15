package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// https://gin-gonic.com/ko-kr/docs/examples/grouping-routes/
func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		/*
			파싱 X, curl -X GET -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bind_json" -i
			파싱 X, curl "localhost:8080/v1/bind_json?price=10" -i
			파싱 O, curl -X POST -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bind_json" -i
			파싱 O(그러나 form 으로
		*/
		v1.GET("/bind_json", bind_json)
		v1.POST("/bind_json", bind_json)

		/*
			파싱 X, curl -X GET -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bind_form" -i
			파싱 O, curl "localhost:8080/v1/bind_json?price=10" -i
			파싱 O, curl -X POST -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bind_form" -i
		*/
		v1.GET("/bind_form", bind_form)
		v1.POST("/bind_form", bind_form)

		/*
			파싱 X, curl "localhost:8080/v1/bindquery_json?price=20" -i
			파싱 X, curl "localhost:8080/v1/bindquery_json?price=10" -i
			파싱 O(10), curl -X POST -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bindquery_json" -i
			파싱 O(그러나 query param이 아닌 body, 20), curl -X POST -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bindquery_json?price=20" -i
		*/
		v1.GET("/bindquery_json", bindquery_json)
		v1.POST("/bindquery_json", bindquery_json)

		/*
			파싱 O, curl "localhost:8080/v1/bindquery_form?price=10" -i
			파싱 O(GET의 body는 작동하지 않고 쿼리 파라미터 20)curl -X GET -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bindquery_form?price=20" -i
			파싱 X, curl -X GET -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bindquery_form" -i
			파싱 O, curl "localhost:8080/v1/bindquery_form?price=20" -i
			파싱 O(그러나 query param이 아니라 Body로 작동) curl -X POST -H "Content-Type: application/json" -d '{"price": 10}' "localhost:8080/v1/bindquery_form?price=20" -i
			파싱 X(파싱안됨.. POST, body가 있는 경우) curl -X POST -H "Content-Type: application/json" -d '{"error": 10}' "localhost:8080/v1/bindquery_form?price=20" -i
			파싱 O(이건됨.. GET, body가 있는 경우) curl -X GET -H "Content-Type: application/json" -d '{"error": 10}' "localhost:8080/v1/bindquery_form?price=20" -i
			파싱 O(파싱됨, POST, 근데 application/json 아닌 경우만) curl -X POST "localhost:8080/v1/bindquery_form?price=20" -i
		*/
		v1.GET("/bindquery_form", bindquery_form)
		v1.POST("/bindquery_form", bindquery_form)

		v1.GET("/shouldbind_json", shouldbind_json)
		v1.POST("/shouldbind_json", shouldbind_json)

	}

	router.Run(":8080")
}

type DataJSON struct {
	Price int `json:"price"`
}

type DataForm struct {
	Price int `form:"price"`
}

func bind_json(c *gin.Context) {
	var testData DataJSON

	err := c.Bind(&testData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "StatusBadRequest")
	}
	c.IndentedJSON(http.StatusOK, testData)
}

func bind_form(c *gin.Context) {
	var testData DataForm

	err := c.Bind(&testData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "StatusBadRequest")
	}
	c.IndentedJSON(http.StatusOK, testData)
}

func bindquery_json(c *gin.Context) {
	var testData DataJSON

	err := c.BindQuery(&testData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "StatusBadRequest")
	}
	c.IndentedJSON(http.StatusOK, testData)
}

func bindquery_form(c *gin.Context) {
	var testData DataForm

	err := c.BindQuery(&testData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "StatusBadRequest")
	}
	c.IndentedJSON(http.StatusOK, testData)
}

func shouldbind_json(c *gin.Context) {
	var testData DataJSON

	err := c.ShouldBind(&testData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "StatusBadRequest")
	}
	c.IndentedJSON(http.StatusOK, testData)
}

func shouldbind_form(c *gin.Context) {
	var testData DataForm

	err := c.ShouldBind(&testData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "StatusBadRequest")
	}
	c.IndentedJSON(http.StatusOK, testData)
}

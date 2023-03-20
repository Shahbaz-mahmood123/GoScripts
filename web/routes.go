package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/orgs", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	router.Run("localhost:8080")
}

func getOrgnizationsRoute() string {

	return "hellow World"
}

type organization struct {
}

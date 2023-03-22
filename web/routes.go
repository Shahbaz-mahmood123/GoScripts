package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	createOrgRoute(router)

	router.GET("/orgs", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	router.Run("localhost:8080")
}

package api

import (
	"fmt"
	"net/http"

	"github.com/Shahbaz-mahmood123/GoScripts/tower"
	"github.com/gin-gonic/gin"
)

func createWorkspaceRoute(router *gin.Engine) {
	router.POST("/orgs/workspace", createWorkspaceHandler)
}

func createWorkspaceHandler(c *gin.Context) {
	var workspace tower.Workspace

	//os.Setenv("LASTCREATEDORG", "174618892238392")

	fmt.Println(c.Request.Body)
	if err := c.BindJSON(&workspace); err != nil {
		fmt.Printf("Error binding JSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	response, err := tower.CreateWorkspace(workspace.Workspace.Name, workspace.Workspace.FullName, workspace.Workspace.Visibility)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating organization"})
		return
	}

	if err != nil {
		panic("something went wrong" + err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{"response": response})
}

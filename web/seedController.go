package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Shahbaz-mahmood123/GoScripts/tower"
	"github.com/gin-gonic/gin"
)

func createSeedDbRoute(router *gin.Engine) {
	router.POST("/seedDatabase", (createSeedDbHandler))
}

func createSeedDbHandler(c *gin.Context) {

	var combinedPayload tower.CombinedPayload
	var responseOrg tower.ReponseOrg

	if err := c.BindJSON(&combinedPayload); err != nil {
		fmt.Printf("Error binding JSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	uri := "orgs"
	orgnizationResponse, orgnizationError := tower.CreateOrganization(uri, combinedPayload.Organization.Name, combinedPayload.Organization.FullName)

	if orgnizationError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating organization"})
		return
	}
	decodedResp := json.Unmarshal([]byte(orgnizationResponse), &responseOrg)
	if decodedResp != nil {
		panic("something went wrong" + decodedResp.Error())
	}

	//set enviornment variable to create workspace in the same org in the workspace controller
	os.Setenv("LASTCREATEDORG", strconv.FormatInt(responseOrg.Organization.OrgID, 10))

	workspaceResponse, workspaceError := tower.CreateWorkspace(combinedPayload.Workspace.Name, combinedPayload.Workspace.FullName, combinedPayload.Workspace.Visibility)

	if workspaceError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating organization"})
		return
	}

	currentTime := time.Now()
	c.JSON(http.StatusCreated, gin.H{"Created": currentTime.String(), "organisation": orgnizationResponse, "workspace": workspaceResponse})

}

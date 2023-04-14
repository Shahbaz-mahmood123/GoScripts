package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"os"
	"strconv"

	"github.com/Shahbaz-mahmood123/GoScripts/tower"
	"github.com/gin-gonic/gin"
)

func createOrgRoute(router *gin.Engine) {
	router.POST("/orgs", (createOrganizationHandler))
}

func createOrganizationHandler(c *gin.Context) {

	// bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	// // Restore the request body data
	// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// // Print the payloadß
	// fmt.Prin(string(bodyBytes))tln
	//Anything below this is not the print code
	var org tower.OrgStruct
	var responseOrg tower.ReponseOrg

	if err := c.BindJSON(&org); err != nil {
		fmt.Printf("Error binding JSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	//fmt.Printf("Received org: %+v\n", org)

	uri := "orgs"
	resp, err := tower.CreateOrganization(uri, org.Organization.Name, org.Organization.FullName)

	decodedResp := json.Unmarshal([]byte(resp), &responseOrg)
	if decodedResp != nil {
		panic("something went wrong" + decodedResp.Error())
	}

	//set enviornment variable to create workspace in the same org
	os.Setenv("LASTCREATEDORG", strconv.FormatInt(responseOrg.Organization.OrgID, 10))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating organization"})
		return
	}

	if err != nil {
		panic("something went wrong" + err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{"response": resp})
}

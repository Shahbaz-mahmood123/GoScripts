package api

import (
	"fmt"
	"net/http"

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
	// // Print the payload
	// fmt.Prin(string(bodyBytes))tln
	//Anything below this is not the print code
	var org tower.OrgStruct

	if err := c.BindJSON(&org); err != nil {
		fmt.Printf("Error binding JSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	//fmt.Printf("Received org: %+v\n", org)

	uri := "orgs"
	resp, err := tower.CreateOrganization(uri, org.Organization.Name, org.Organization.FullName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating organization"})
		return
	}

	if err != nil {
		panic("something went wrong" + err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{"response": resp})
}

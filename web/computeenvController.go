package api

import (
	"fmt"
	"net/http"

	"github.com/Shahbaz-mahmood123/GoScripts/tower"
	"github.com/gin-gonic/gin"
)

func createComputeEnvRoute(router *gin.Engine) {
	router.POST("/compute", (createComputeEnvHandler))
}

func createComputeEnvHandler(c *gin.Context) {

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
	var compute tower.OrgStruct

	if err := c.BindJSON(&compute); err != nil {
		fmt.Printf("Error binding JSON: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	//fmt.Printf("Received org: %+v\n", org)

	uri := "compute-envs?workspaceId=251085962711837"
	resp, err := tower.CreateComputeEnv(uri)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating organization"})
		return
	}

	if err != nil {
		panic("something went wrong" + err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{"response": resp})
}

package main

import (
	"fmt"

	"github.com/Shahbaz-mahmood123/GoScripts/tower"
	api "github.com/Shahbaz-mahmood123/GoScripts/web"
)

func main() {

	//tower.GetComputeEnv()
	response := tower.GetOrganizations()
	api.Routes()
	fmt.Println(response)
}

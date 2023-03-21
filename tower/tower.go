package tower

import (
	"fmt"
)

func GetComputeEnv() {
	uri := "compute-envs?workspaceId=251085962711837"
	response, err := httpClient(uri, "GET", nil)

	if err != nil {
		panic(err)
	}
	fmt.Println(response)

}

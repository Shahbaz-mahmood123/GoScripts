package tower

import (
	"fmt"
)

func GetComputeEnv() {
	uri := "compute-envs?workspaceId=251085962711837"
	response := httpClient(uri)

	fmt.Println(response)

}

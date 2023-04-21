package tower

import "time"

func fetchComputeEnvs() string {
	uri := "compute-envs?workspaceId=251085962711837"

	response, err := httpClient(uri, GET, nil)
	if err != nil {
		panic(err)
	}
	return response
}

type Compute struct {
	ComputeEnvs []struct {
		ID            string    `json:"id"`
		Name          string    `json:"name"`
		Platform      string    `json:"platform"`
		Status        string    `json:"status"`
		Message       any       `json:"message"`
		LastUsed      time.Time `json:"lastUsed"`
		Primary       any       `json:"primary"`
		WorkspaceName string    `json:"workspaceName"`
		Visibility    string    `json:"visibility"`
		WorkDir       string    `json:"workDir"`
	} `json:"computeEnvs"`
}

func CreateComputeEnv(uri string) (string, error) {

	return "", nil
}

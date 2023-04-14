package tower

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetWorkspace() {

}

type Workspace struct {
	WorkspaceData struct {
		Name       string `json:"name"`
		FullName   string `json:"fullName"`
		Visibility string `json:"visibility"`
	} `json:"workspace"`
}

func CreateWorkspace(name string, fullName string, visibility string) (string, error) {

	lastCreatedOrg := os.Getenv("LASTCREATEDORG")
	uri := "orgs/" + lastCreatedOrg + "/workspaces"

	createWorkspaceRequest := Workspace{
		WorkspaceData: struct {
			Name       string `json:"name"`
			FullName   string `json:"fullName"`
			Visibility string `json:"visibility"`
		}{
			Name:       name,
			FullName:   fullName,
			Visibility: visibility,
		},
	}

	payload, err := json.Marshal(createWorkspaceRequest)
	if err != nil {
		return "", err
	}

	resp, err := httpClient(uri, http.MethodPost, payload)
	if err != nil {
		return "", err
	}

	return resp, nil
}

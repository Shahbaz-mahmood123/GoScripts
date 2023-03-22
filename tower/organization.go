package tower

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetOrganizations() string {

	uri := "orgs"

	response, err := httpClient(uri, GET, nil)
	if err != nil {
		panic(err)
	}
	return response
}

type OrgStruct struct {
	Organization struct {
		Name     string `json:"name"`
		FullName string `json:"fullName"`
	} `json:"organization"`
}

func CreateOrganization(uri string, name string, fullName string) (string, error) {

	createOrgRequest := OrgStruct{
		Organization: struct {
			Name     string `json:"name"`
			FullName string `json:"fullName"`
		}{
			Name:     name,
			FullName: fullName,
		},
	}

	fmt.Println(createOrgRequest)
	payload, err := json.Marshal(createOrgRequest)
	if err != nil {
		return "", err
	}

	resp, err := httpClient(uri, http.MethodPost, payload)
	if err != nil {
		return "", err
	}

	return resp, nil
}

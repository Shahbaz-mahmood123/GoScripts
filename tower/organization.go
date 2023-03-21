package tower

func GetOrganizations() string {
	uri := "orgs"

	response, err := httpClient(uri, "GET", nil)
	if err != nil {
		panic(err)
	}
	return response
}

func CreateOrganizations() string {

	return ""
}

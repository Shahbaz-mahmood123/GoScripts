package tower

func GetOrganizations() string {
	uri := "orgs"

	response, err := httpClient(uri)
	if err != nil {
		panic(err)
	}
	return response
}

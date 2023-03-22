package tower

func GetServiceInfo() string {
	uri := "service-info"

	response, err := httpClient(uri, GET, nil)
	if err != nil {
		panic(err)
	}
	return response

}

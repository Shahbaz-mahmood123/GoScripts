package tower

func GetServiceInfo() string {
	uri := "service-info"

	response, err := httpClient(uri)
	if err != nil {
		panic(err)
	}
	return response

}

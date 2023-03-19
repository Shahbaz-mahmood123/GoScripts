package tower

import "fmt"

func getServiceInfo() {
	uri := "service-info"

	response := httpClient(uri)

	fmt.Println(response)

}
